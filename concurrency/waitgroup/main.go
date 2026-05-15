package main

import (
	"fmt"
	"sync"
	"time"
)

/*
  WAITGROUP IN GO

KEY POINTS:
  - WaitGroup waits for a collection of goroutines to finish
  - Part of the 'sync' package
  - Better than time.Sleep() - waits exactly as long as needed
  - Counter-based: Add() increases, Done() decreases, Wait() blocks until 0

METHODS:
  wg.Add(n)  - Add n goroutines to wait for
  wg.Done()  - Mark one goroutine as complete (decrements counter)
  wg.Wait()  - Block until counter reaches 0

COMMON PATTERNS:
  1. Add before starting goroutine
  2. defer wg.Done() at start of goroutine
  3. Wait after all goroutines started

CAUTION:
  - Don't copy WaitGroup after first use (pass by pointer)
  - Add() must be called before Wait()
  - Negative counter causes panic

*/

func main() {
	// 1. BASIC WAITGROUP

	fmt.Println("--- Basic WaitGroup ---")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment counter BEFORE goroutine
		go func(n int) {
			defer wg.Done() // decrement when done
			fmt.Printf("Worker %d starting\n", n)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Worker %d done\n", n)
		}(i)
	}

	wg.Wait() // blocks until counter is 0
	fmt.Println("All workers completed")

	// 2. WAITGROUP WITH NAMED FUNCTION

	fmt.Println("\n--- WaitGroup with Function ---")

	var wg2 sync.WaitGroup
	tasks := []string{"task-A", "task-B", "task-C"}

	for _, task := range tasks {
		wg2.Add(1)
		go worker(task, &wg2) // pass WaitGroup by pointer
	}

	wg2.Wait()
	fmt.Println("All tasks completed")

	// 3. BATCH PROCESSING

	fmt.Println("\n--- Batch Processing ---")

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg3 sync.WaitGroup

	for _, item := range items {
		wg3.Add(1)
		go func(n int) {
			defer wg3.Done()
			result := process(n)
			fmt.Printf("Processed %d -> %d\n", n, result)
		}(item)
	}

	wg3.Wait()
	fmt.Println("Batch complete")

	// 4. COLLECTING RESULTS (WITH MUTEX)

	fmt.Println("\n--- Collecting Results ---")

	var wg4 sync.WaitGroup
	var mu sync.Mutex
	results := make([]int, 0)

	for i := 1; i <= 5; i++ {
		wg4.Add(1)
		go func(n int) {
			defer wg4.Done()
			result := n * n

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}(i)
	}

	wg4.Wait()
	fmt.Println("Results:", results)

	// 5. WORKER POOL PATTERN

	fmt.Println("\n--- Worker Pool ---")

	jobs := make(chan int, 10)
	var wg5 sync.WaitGroup

	// Start 3 workers
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		wg5.Add(1)
		go poolWorker(w, jobs, &wg5)
	}

	// Send jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs) // signal no more jobs

	wg5.Wait()
	fmt.Println("All jobs completed")

	// 6. NESTED WAITGROUPS

	fmt.Println("\n--- Nested WaitGroups ---")

	var outer sync.WaitGroup

	for i := 1; i <= 2; i++ {
		outer.Add(1)
		go func(groupID int) {
			defer outer.Done()

			var inner sync.WaitGroup
			for j := 1; j <= 3; j++ {
				inner.Add(1)
				go func(taskID int) {
					defer inner.Done()
					time.Sleep(50 * time.Millisecond)
					fmt.Printf("Group %d, Task %d done\n", groupID, taskID)
				}(j)
			}
			inner.Wait()
			fmt.Printf("Group %d complete\n", groupID)
		}(i)
	}

	outer.Wait()
	fmt.Println("All groups complete")

	// 7. TIMEOUT PATTERN

	fmt.Println("\n--- Timeout Pattern ---")

	var wg6 sync.WaitGroup
	done := make(chan struct{})

	wg6.Add(1)
	go func() {
		defer wg6.Done()
		time.Sleep(200 * time.Millisecond) // slow task
		fmt.Println("Slow task completed")
	}()

	// Wait in separate goroutine
	go func() {
		wg6.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Tasks finished in time")
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Timeout waiting for tasks")
	}

	// 8. ERROR HANDLING PATTERN

	fmt.Println("\n--- Error Handling ---")

	var wg7 sync.WaitGroup
	errChan := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		wg7.Add(1)
		go func(n int) {
			defer wg7.Done()
			if err := riskyOperation(n); err != nil {
				errChan <- err
			}
		}(i)
	}

	wg7.Wait()
	close(errChan)

	// Collect errors
	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		fmt.Println("Errors occurred:", errors)
	} else {
		fmt.Println("All operations succeeded")
	}

	// 9. PERFORMANCE NOTE

	fmt.Println("\n--- Performance Note ---")
	fmt.Println(`
Remember: More goroutines != always faster
- Each goroutine has overhead (scheduling, memory)
- For CPU-bound tasks, optimal = GOMAXPROCS goroutines
- For I/O-bound tasks, more goroutines can help
- Always benchmark your specific use case
`)
}

// HELPER FUNCTIONS

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s: starting\n", name)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("%s: done\n", name)
}

func process(n int) int {
	time.Sleep(50 * time.Millisecond)
	return n * n
}

func poolWorker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("Worker %d finished\n", id)
}

func riskyOperation(n int) error {
	time.Sleep(50 * time.Millisecond)
	if n == 2 {
		return fmt.Errorf("operation %d failed", n)
	}
	return nil
}
