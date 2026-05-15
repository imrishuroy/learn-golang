package main

import (
	"fmt"
	"os"
	"time"
)

/*
  DEFER, PANIC, AND RECOVER IN GO

KEY POINTS:
  - defer: Schedule function call to run when surrounding function returns
  - panic: Stop normal execution and begin unwinding the stack
  - recover: Regain control after a panic

DEFER:
  - Deferred calls are executed in LIFO order (stack)
  - Arguments are evaluated when defer is executed, not when function runs
  - Commonly used for cleanup: closing files, releasing locks, timing

PANIC:
  - Caused by runtime errors or explicit panic() call
  - Unwinds the stack, running deferred functions along the way
  - Program crashes if panic reaches top of goroutine stack

RECOVER:
  - Only works inside deferred functions
  - Returns nil if no panic is occurring
  - Used to handle panics gracefully

*/

func main() {
	// 1. DEFER BASICS

	fmt.Println("--- Defer Basics ---")

	fmt.Println("Start")
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")
	fmt.Println("End")
	fmt.Println("(Deferred calls run after End, in reverse order)")
	fmt.Println()

	// 2. DEFER STACK (LIFO)

	fmt.Println("--- Defer Stack (LIFO) ---")
	countDown()

	// 3. DEFER ARGUMENT EVALUATION

	fmt.Println("\n--- Defer Argument Evaluation ---")

	x := 10
	defer fmt.Printf("Deferred with arg (captured at defer): x=%d\n", x)

	defer func() {
		fmt.Printf("Deferred closure (sees final value): x=%d\n", x)
	}()

	x = 20
	fmt.Printf("Current x: %d\n", x)

	// 4. PRACTICAL: TIMING FUNCTIONS

	fmt.Println("\n--- Practical: Function Timing ---")

	slowOperation()

	// 5. PRACTICAL: RESOURCE CLEANUP

	fmt.Println("\n--- Practical: File Cleanup Pattern ---")

	// Simulated file handling
	processFileDemo()

	// 6. PANIC BASICS

	fmt.Println("\n--- Panic Basics ---")

	// This would panic and crash:
	// panic("something went wrong!")

	// Common runtime panics:
	// - Array/slice index out of bounds
	// - Nil pointer dereference
	// - Type assertion failure

	fmt.Println("Panic can be caused by:")
	fmt.Println("  - panic(\"message\")")
	fmt.Println("  - Array index out of bounds")
	fmt.Println("  - Nil pointer dereference")
	fmt.Println("  - Failed type assertion")

	// 7. RECOVER FROM PANIC

	fmt.Println("\n--- Recover from Panic ---")

	// Safe function that recovers from panic
	result := safeOperation(func() {
		fmt.Println("About to panic...")
		panic("oops!")
	})
	fmt.Println("Recovered:", result)

	// Normal operation continues
	result = safeOperation(func() {
		fmt.Println("Normal operation")
	})
	fmt.Println("Normal result:", result)

	// 8. RECOVER PATTERN

	fmt.Println("\n--- Recover Pattern ---")

	// Calling function that might panic
	safeDivide(10, 2)
	safeDivide(10, 0) // would panic without recover
	safeDivide(20, 4)

	// 9. MULTIPLE DEFERS WITH PANIC

	fmt.Println("\n--- Defers Run During Panic ---")

	func() {
		defer fmt.Println("  Deferred 1 (runs)")
		defer fmt.Println("  Deferred 2 (runs)")
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("  Recovered:", r)
			}
		}()
		defer fmt.Println("  Deferred 3 (runs)")

		fmt.Println("  About to panic...")
		panic("test panic")
		fmt.Println("  This never prints") // unreachable
	}()

	fmt.Println("Execution continues after recovery")

	// 10. PRACTICAL: GRACEFUL SHUTDOWN

	fmt.Println("\n--- Practical: Graceful Shutdown ---")

	runServer()

	// 11. DEFER IN LOOPS

	fmt.Println("\n--- Defer in Loops (Careful!) ---")

	// BAD: Defers accumulate, run after loop finishes
	fmt.Println("BAD pattern (avoid):")
	badLoopDefer()

	// GOOD: Use wrapper function
	fmt.Println("\nGOOD pattern:")
	goodLoopDefer()

	// 12. PRACTICAL: MUTEX UNLOCK

	fmt.Println("\n--- Practical: Lock/Unlock Pattern ---")

	// Simulated critical section
	criticalSection()

	// 13. CONVERTING PANIC TO ERROR

	fmt.Println("\n--- Convert Panic to Error ---")

	err := riskyOperationWrapper()
	if err != nil {
		fmt.Println("Caught error:", err)
	}

	// 14. NAMED RETURNS WITH DEFER

	fmt.Println("\n--- Named Returns with Defer ---")

	result1 := namedReturnDemo(1)
	result2 := namedReturnDemo(2)
	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
}

// HELPER FUNCTIONS

func countDown() {
	for i := 5; i >= 1; i-- {
		defer fmt.Printf("%d ", i)
	}
	fmt.Print("Countdown: ")
}

func slowOperation() {
	start := time.Now()
	defer func() {
		fmt.Printf("slowOperation took %v\n", time.Since(start))
	}()

	// Simulate work
	time.Sleep(50 * time.Millisecond)
}

func processFileDemo() {
	fmt.Println("Opening file...")

	// In real code: file, err := os.Open("...")
	defer func() {
		fmt.Println("Closing file... (cleanup)")
		// In real code: file.Close()
	}()

	fmt.Println("Reading file...")
	fmt.Println("Processing data...")
}

func safeOperation(fn func()) (err string) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Sprintf("panic recovered: %v", r)
		}
	}()

	fn()
	return "success"
}

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("safeDivide: recovered from %v\n", r)
		}
	}()

	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func runServer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Server crashed, performing cleanup...")
			fmt.Println("  - Closing connections")
			fmt.Println("  - Saving state")
			fmt.Println("  - Logging error:", r)
		}
	}()

	fmt.Println("Server starting...")
	fmt.Println("Server running... (simulated crash)")
	panic("server error")
}

func badLoopDefer() {
	for i := 0; i < 3; i++ {
		// All 3 defers run AFTER loop completes
		defer fmt.Printf("  deferred: %d\n", i)
	}
	fmt.Println("  Loop finished")
}

func goodLoopDefer() {
	for i := 0; i < 3; i++ {
		func(n int) {
			defer fmt.Printf("  deferred: %d\n", n)
			// Each defer runs when this inner function returns
		}(i)
	}
	fmt.Println("  Loop finished")
}

func criticalSection() {
	// Simulating mutex lock/unlock
	fmt.Println("Acquiring lock...")
	// mu.Lock()

	defer func() {
		fmt.Println("Releasing lock...")
		// mu.Unlock()
	}()

	fmt.Println("Doing critical work...")

	// Even if we return early or panic, unlock happens
	if true {
		fmt.Println("Early return")
		return
	}
}

func riskyOperationWrapper() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("operation failed: %v", r)
		}
	}()

	// Simulate risky operation
	riskyOperation()
	return nil
}

func riskyOperation() {
	panic("something went wrong")
}

func namedReturnDemo(x int) (result int) {
	result = x

	defer func() {
		result *= 2 // modify named return
	}()

	return result // original value modified by defer
}

// REAL-WORLD FILE HANDLING EXAMPLE

func processRealFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // guaranteed cleanup

	// Process file...
	// Even if this panics, file.Close() runs

	return nil
}
