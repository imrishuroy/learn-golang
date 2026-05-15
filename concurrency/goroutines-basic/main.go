package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
  GOROUTINES IN GO

KEY POINTS:
  - Goroutines are lightweight threads managed by Go runtime
  - Much cheaper than OS threads (~2KB stack vs ~1MB)
  - Created with the 'go' keyword before a function call
  - Main goroutine exits = all other goroutines die
  - Go scheduler multiplexes goroutines onto OS threads

SYNTAX:
  go functionName(args)           // start goroutine
  go func() { ... }()             // anonymous goroutine

IMPORTANT:
  - Goroutines run CONCURRENTLY, not necessarily in PARALLEL
  - Concurrent = dealing with multiple things at once (structure)
  - Parallel = doing multiple things at once (execution)
  - No guaranteed execution order
  - Need synchronization to communicate (channels, WaitGroup, mutex)

*/

func main() {
	// 1. BASIC GOROUTINE

	fmt.Println("--- Basic Goroutine ---")

	// Start a goroutine
	go sayHello("goroutine")

	// Main continues immediately
	fmt.Println("Main: started goroutine")

	// Without this sleep, main exits before goroutine runs!
	time.Sleep(100 * time.Millisecond)

	// 2. MULTIPLE GOROUTINES

	fmt.Println("\n--- Multiple Goroutines ---")

	for i := 1; i <= 3; i++ {
		go func(n int) {
			fmt.Printf("Goroutine %d started\n", n)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Goroutine %d finished\n", n)
		}(i)
	}

	time.Sleep(200 * time.Millisecond)

	// 3. GOROUTINE CLOSURE GOTCHA

	fmt.Println("\n--- Closure Gotcha ---")

	// WRONG: All goroutines see the same 'i'
	fmt.Print("Wrong way: ")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i, " ") // likely prints 5 5 5 5 5
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// CORRECT: Pass 'i' as parameter
	fmt.Print("Right way: ")
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Print(n, " ") // prints 0-4 in some order
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// 4. ANONYMOUS VS NAMED FUNCTION

	fmt.Println("\n--- Anonymous vs Named ---")

	// Named function
	go printNumbers("named", 3)

	// Anonymous function
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("anonymous: %d\n", i)
			time.Sleep(30 * time.Millisecond)
		}
	}()

	time.Sleep(200 * time.Millisecond)

	// 5. RUNTIME INFO

	fmt.Println("\n--- Runtime Info ---")

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // current value
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	// Start some goroutines
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Second)
		}()
	}

	fmt.Println("After starting 10 goroutines:", runtime.NumGoroutine())

	// 6. GOROUTINES ARE CHEAP

	fmt.Println("\n--- Goroutines are Cheap ---")

	start := time.Now()

	// Create many goroutines
	count := 10000
	for i := 0; i < count; i++ {
		go func() {
			time.Sleep(10 * time.Millisecond)
		}()
	}

	fmt.Printf("Created %d goroutines in %v\n", count, time.Since(start))
	fmt.Println("Active goroutines:", runtime.NumGoroutine())

	// 7. FIRE AND FORGET PATTERN

	fmt.Println("\n--- Fire and Forget ---")

	// Background task that doesn't need result
	go logToFile("User logged in")
	go sendEmail("user@example.com")

	fmt.Println("Tasks dispatched")
	time.Sleep(100 * time.Millisecond)

	// 8. CONCURRENT EXECUTION DEMO

	fmt.Println("\n--- Concurrent Execution ---")

	// Sequential (slow)
	start = time.Now()
	result1 := slowOperation("A")
	result2 := slowOperation("B")
	result3 := slowOperation("C")
	fmt.Printf("Sequential: %s %s %s in %v\n", result1, result2, result3, time.Since(start))

	// Concurrent (fast) - but we need channels/waitgroup to get results properly
	// This is just a demo - results won't be captured correctly
	start = time.Now()
	go slowOperation("X")
	go slowOperation("Y")
	go slowOperation("Z")
	time.Sleep(150 * time.Millisecond) // wait for completion
	fmt.Printf("Concurrent: completed in ~%v\n", time.Since(start))

	fmt.Println("\n--- End of Demo ---")
	fmt.Println("Note: Use WaitGroup or Channels for proper synchronization!")
	fmt.Println("See: 04-concurrency/waitgroup and 04-concurrency/channels")
}

// HELPER FUNCTIONS

func sayHello(from string) {
	fmt.Printf("Hello from %s!\n", from)
}

func printNumbers(name string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(30 * time.Millisecond)
	}
}

func logToFile(message string) {
	time.Sleep(20 * time.Millisecond) // simulate I/O
	fmt.Println("[LOG]", message)
}

func sendEmail(to string) {
	time.Sleep(30 * time.Millisecond) // simulate network
	fmt.Println("[EMAIL] Sent to", to)
}

func slowOperation(name string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("Result-%s", name)
}
