package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
  CONTEXT PACKAGE IN GO

KEY POINTS:
  - Context carries deadlines, cancellation signals, and request-scoped values
  - Essential for controlling goroutine lifecycles
  - Should be first parameter of functions: func DoSomething(ctx context.Context, ...)
  - Never store Context in a struct; pass it explicitly

TYPES OF CONTEXT:
  context.Background()       // Root context, never cancelled
  context.TODO()            // Placeholder when unsure which context to use
  context.WithCancel()      // Cancellable context
  context.WithTimeout()     // Auto-cancels after duration
  context.WithDeadline()    // Auto-cancels at specific time
  context.WithValue()       // Carries request-scoped data

USE CASES:
  - HTTP request handling (request cancellation)
  - Database queries with timeout
  - API calls with deadline
  - Graceful shutdown
  - Propagating cancellation across goroutines

*/

func main() {
	// 1. BACKGROUND AND TODO CONTEXT

	fmt.Println("--- Background and TODO ---")

	// Background: root of any context tree
	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)

	// TODO: use when unsure, placeholder
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n", todoCtx)

	// 2. CONTEXT WITH CANCEL

	fmt.Println("\n--- WithCancel ---")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker: received cancellation")
				return
			default:
				fmt.Println("Worker: working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(350 * time.Millisecond)
	fmt.Println("Main: cancelling worker")
	cancel() // Signal cancellation
	time.Sleep(50 * time.Millisecond)

	// 3. CONTEXT WITH TIMEOUT

	fmt.Println("\n--- WithTimeout ---")

	ctx, cancel = context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel() // Always call cancel to release resources

	fmt.Println("Starting slow operation with 200ms timeout")

	select {
	case result := <-slowOperation(ctx):
		fmt.Println("Result:", result)
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}

	// 4. CONTEXT WITH DEADLINE

	fmt.Println("\n--- WithDeadline ---")

	deadline := time.Now().Add(150 * time.Millisecond)
	ctx, cancel = context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("Deadline set to: %v\n", deadline.Format("15:04:05.000"))

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Deadline exceeded:", ctx.Err())
	}

	// 5. CONTEXT WITH VALUES

	fmt.Println("\n--- WithValue ---")

	// Create context with values
	ctx = context.WithValue(context.Background(), "userID", 42)
	ctx = context.WithValue(ctx, "requestID", "req-12345")

	// Pass to function
	handleRequest(ctx)

	// 6. PROPAGATING CANCELLATION

	fmt.Println("\n--- Propagating Cancellation ---")

	ctx, cancel = context.WithCancel(context.Background())

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	time.Sleep(250 * time.Millisecond)
	fmt.Println("Cancelling all workers...")
	cancel()
	wg.Wait()
	fmt.Println("All workers stopped")

	// 7. CHECKING CONTEXT STATE

	fmt.Println("\n--- Checking Context State ---")

	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)

	// Check deadline
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Println("Deadline:", deadline.Format("15:04:05.000"))
	}

	// Check if done (not yet)
	select {
	case <-ctx.Done():
		fmt.Println("Already done")
	default:
		fmt.Println("Not done yet")
	}

	time.Sleep(150 * time.Millisecond)

	// Now it's done
	fmt.Println("After timeout - Err:", ctx.Err())
	cancel()

	// 8. PRACTICAL: HTTP-STYLE REQUEST HANDLING

	fmt.Println("\n--- Practical: Request Handling ---")

	ctx, cancel = context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	result, err := fetchDataWithContext(ctx, "https://api.example.com/data")
	if err != nil {
		fmt.Println("Fetch error:", err)
	} else {
		fmt.Println("Fetch result:", result)
	}

	// 9. PRACTICAL: DATABASE QUERY SIMULATION

	fmt.Println("\n--- Practical: Database Query ---")

	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err = queryDatabase(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Println("Query error:", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	users, err := queryDatabase(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Println("Query error:", err)
	} else {
		fmt.Println("Query result:", users)
	}

	// 10. GRACEFUL SHUTDOWN PATTERN

	fmt.Println("\n--- Graceful Shutdown ---")

	runServer()

	// 11. CONTEXT BEST PRACTICES

	fmt.Println("\n--- Best Practices ---")

	fmt.Println("1. Pass context as first parameter")
	fmt.Println("2. Don't store context in structs")
	fmt.Println("3. Always call cancel (use defer)")
	fmt.Println("4. Don't pass nil context, use context.TODO()")
	fmt.Println("5. Only use WithValue for request-scoped data")
	fmt.Println("6. Keep context values to a minimum")

	// 12. CONTEXT INHERITANCE

	fmt.Println("\n--- Context Inheritance ---")

	parentCtx, parentCancel := context.WithCancel(context.Background())

	childCtx1, childCancel1 := context.WithTimeout(parentCtx, 500*time.Millisecond)
	defer childCancel1()

	childCtx2, childCancel2 := context.WithCancel(parentCtx)
	defer childCancel2()

	var wg2 sync.WaitGroup
	wg2.Add(2)

	go func() {
		defer wg2.Done()
		<-childCtx1.Done()
		fmt.Println("Child 1 cancelled:", childCtx1.Err())
	}()

	go func() {
		defer wg2.Done()
		<-childCtx2.Done()
		fmt.Println("Child 2 cancelled:", childCtx2.Err())
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Cancelling parent (cancels all children)")
	parentCancel()

	wg2.Wait()
}

// HELPER FUNCTIONS

func slowOperation(ctx context.Context) <-chan string {
	result := make(chan string)

	go func() {
		// Simulate slow work
		time.Sleep(300 * time.Millisecond)

		select {
		case <-ctx.Done():
			// Context cancelled, don't send result
			return
		case result <- "operation completed":
		}
	}()

	return result
}

func handleRequest(ctx context.Context) {
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")

	fmt.Printf("Handling request %v for user %v\n", requestID, userID)
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d: working\n", id)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func fetchDataWithContext(ctx context.Context, url string) (string, error) {
	// Simulate HTTP request
	resultCh := make(chan string)
	errCh := make(chan error)

	go func() {
		// Simulate network delay
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
		resultCh <- fmt.Sprintf("data from %s", url)
	}()

	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func queryDatabase(ctx context.Context, query string) ([]string, error) {
	resultCh := make(chan []string)

	go func() {
		// Simulate slow query
		time.Sleep(200 * time.Millisecond)
		resultCh <- []string{"user1", "user2", "user3"}
	}()

	select {
	case result := <-resultCh:
		return result, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("query cancelled: %w", ctx.Err())
	}
}

func runServer() {
	ctx, cancel := context.WithCancel(context.Background())

	// Simulated server
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Server: starting...")

		<-ctx.Done()
		fmt.Println("Server: shutting down gracefully...")
		time.Sleep(50 * time.Millisecond) // Cleanup time
		fmt.Println("Server: stopped")
	}()

	// Simulate running for a bit then shutdown
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Signal: initiating shutdown")
	cancel()

	wg.Wait()
}
