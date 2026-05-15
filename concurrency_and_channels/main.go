package main

import (
	"fmt"
	"sync"
	"time"
)

/*
  CONCURRENCY IN GO: GOROUTINES, CHANNELS, AND SYNC

KEY POINTS:
  - Goroutines: Lightweight threads managed by Go runtime
  - Channels: Typed conduits for goroutine communication
  - WaitGroup: Wait for goroutines to complete
  - Mutex: Protect shared data from race conditions
  - Select: Handle multiple channel operations

"Don't communicate by sharing memory; share memory by communicating."

GOROUTINES:
  go functionName()        // Start goroutine
  go func() { ... }()      // Anonymous goroutine

CHANNELS:
  ch := make(chan Type)    // Unbuffered channel
  ch := make(chan Type, n) // Buffered channel (capacity n)
  ch <- value              // Send
  value := <-ch            // Receive
  close(ch)                // Close channel

SYNC PRIMITIVES:
  sync.WaitGroup           // Wait for goroutines
  sync.Mutex               // Mutual exclusion lock
  sync.RWMutex             // Read/Write lock

*/

func main() {
	// 1. BASIC GOROUTINE

	fmt.Println("--- Basic Goroutine ---")

	go sayHello("Goroutine")
	fmt.Println("Main: Started goroutine")
	time.Sleep(100 * time.Millisecond) // Wait for goroutine

	// 2. WAITGROUP - PROPER SYNCHRONIZATION

	fmt.Println("\n--- WaitGroup ---")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", n)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Worker %d done\n", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers finished")

	// 3. UNBUFFERED CHANNELS

	fmt.Println("\n--- Unbuffered Channels ---")

	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine!"
	}()

	msg := <-ch // Blocks until message received
	fmt.Println("Received:", msg)

	// 4. BUFFERED CHANNELS

	fmt.Println("\n--- Buffered Channels ---")

	buffered := make(chan int, 3)

	// Can send without blocking (up to buffer size)
	buffered <- 1
	buffered <- 2
	buffered <- 3

	fmt.Println("Buffer length:", len(buffered))
	fmt.Println("Buffer capacity:", cap(buffered))

	fmt.Println(<-buffered) // 1
	fmt.Println(<-buffered) // 2
	fmt.Println(<-buffered) // 3

	// 5. CHANNEL DIRECTIONS

	fmt.Println("\n--- Channel Directions ---")

	ch2 := make(chan int)

	go sender(ch2)   // Send-only in function
	receiver(ch2)    // Receive-only in function

	// 6. CLOSING CHANNELS AND RANGE

	fmt.Println("\n--- Closing Channels ---")

	numbers := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // Signal no more values
	}()

	// Range stops when channel is closed
	fmt.Print("Numbers: ")
	for num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// Check if channel is closed
	closedCh := make(chan int)
	close(closedCh)
	val, ok := <-closedCh
	fmt.Printf("Closed channel: value=%d, open=%v\n", val, ok)

	// 7. SELECT STATEMENT

	fmt.Println("\n--- Select Statement ---")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		c1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		c2 <- "from channel 2"
	}()

	// Receive from whichever is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		}
	}

	// 8. SELECT WITH TIMEOUT

	fmt.Println("\n--- Select with Timeout ---")

	slowCh := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		slowCh <- "slow response"
	}()

	select {
	case result := <-slowCh:
		fmt.Println("Got:", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	// Drain the channel to avoid goroutine leak
	go func() { <-slowCh }()

	// 9. SELECT WITH DEFAULT (NON-BLOCKING)

	fmt.Println("\n--- Non-blocking Select ---")

	messages := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available")
	}

	// 10. MUTEX - PROTECTING SHARED DATA

	fmt.Println("\n--- Mutex ---")

	var (
		counter int
		mu      sync.Mutex
		wg2     sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg2.Wait()
	fmt.Println("Counter:", counter) // Always 100

	// 11. WORKER POOL PATTERN

	fmt.Println("\n--- Worker Pool ---")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		fmt.Println("Result:", <-results)
	}

	// 12. DONE CHANNEL PATTERN

	fmt.Println("\n--- Done Channel Pattern ---")

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Worker: shutting down")
				return
			default:
				fmt.Println("Worker: working...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(150 * time.Millisecond)
	close(done) // Signal worker to stop
	time.Sleep(50 * time.Millisecond)

	// 13. FAN-OUT FAN-IN

	fmt.Println("\n--- Fan-Out Fan-In ---")

	input := make(chan int)
	output := fanIn(
		square(input),
		square(input),
	)

	go func() {
		for i := 1; i <= 4; i++ {
			input <- i
		}
		close(input)
	}()

	for result := range output {
		fmt.Printf("Squared: %d\n", result)
	}

	fmt.Println("\n--- End ---")
}

// HELPER FUNCTIONS

func sayHello(from string) {
	fmt.Printf("Hello from %s!\n", from)
}

// sender only sends to channel
func sender(ch chan<- int) {
	ch <- 42
}

// receiver only receives from channel
func receiver(ch <-chan int) {
	fmt.Println("Received:", <-ch)
}

// worker processes jobs from channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(20 * time.Millisecond)
		results <- j * 2
	}
}

// square reads from input, squares, sends to output
func square(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for n := range input {
			output <- n * n
		}
		close(output)
	}()
	return output
}

// fanIn merges multiple channels into one
func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	output := make(chan int)

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				output <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
