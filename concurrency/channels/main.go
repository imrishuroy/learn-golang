package main

import (
	"fmt"
	"sync"
	"time"
)

/*
  CHANNELS IN GO

KEY POINTS:
  - Channels are typed conduits for goroutine communication
  - "Don't communicate by sharing memory; share memory by communicating"
  - Channels synchronize execution AND transfer data
  - Zero value of a channel is nil (unusable)

TYPES:
  1. Unbuffered: ch := make(chan T)
     - Synchronous: sender blocks until receiver ready
     - Guarantees delivery before sender continues

  2. Buffered: ch := make(chan T, size)
     - Asynchronous up to buffer size
     - Sender blocks only when buffer is full
     - Receiver blocks only when buffer is empty

OPERATIONS:
  ch <- value    // send
  value := <-ch  // receive
  close(ch)      // close (only sender should close)

DIRECTION:
  chan T         // bidirectional
  chan<- T       // send-only
  <-chan T       // receive-only

*/

func main() {
	// 1. UNBUFFERED CHANNEL - BASIC

	fmt.Println("--- Unbuffered Channel ---")

	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine!"
	}()

	msg := <-ch // blocks until message received
	fmt.Println(msg)

	// 2. UNBUFFERED CHANNEL - SYNCHRONIZATION

	fmt.Println("\n--- Synchronization ---")

	done := make(chan bool)

	go func() {
		fmt.Println("Working...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Done!")
		done <- true
	}()

	<-done // wait for goroutine to finish
	fmt.Println("Main: goroutine completed")

	// 3. BUFFERED CHANNELS

	fmt.Println("\n--- Buffered Channel ---")

	buffered := make(chan int, 3) // buffer size 3

	// Can send without blocking (up to buffer size)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	// buffered <- 4 // would block!

	fmt.Println("Buffer length:", len(buffered))
	fmt.Println("Buffer capacity:", cap(buffered))

	fmt.Println(<-buffered) // 1
	fmt.Println(<-buffered) // 2
	fmt.Println(<-buffered) // 3

	// 4. CHANNEL DIRECTIONS

	fmt.Println("\n--- Channel Directions ---")

	ch2 := make(chan int)

	go send(ch2)    // send-only in function
	receive(ch2)    // receive-only in function

	// 5. CLOSING CHANNELS

	fmt.Println("\n--- Closing Channels ---")

	numbers := make(chan int, 5)

	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // signal no more values
	}()

	// Consumer - range stops when channel closed
	for num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// Check if channel is closed
	closedCh := make(chan int)
	close(closedCh)
	val, ok := <-closedCh
	fmt.Printf("Value: %d, Open: %v\n", val, ok) // 0, false

	// 6. SELECT STATEMENT

	fmt.Println("\n--- Select Statement ---")

	ch1 := make(chan string)
	ch3 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch3 <- "from ch3"
	}()

	// Receive from whichever is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg3 := <-ch3:
			fmt.Println("Received:", msg3)
		}
	}

	// 7. SELECT WITH DEFAULT (NON-BLOCKING)

	fmt.Println("\n--- Non-blocking Select ---")

	messages := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available")
	}

	// Non-blocking send
	select {
	case messages <- "hello":
		fmt.Println("Sent message")
	default:
		fmt.Println("No receiver ready")
	}

	// 8. TIMEOUT WITH SELECT

	fmt.Println("\n--- Timeout ---")

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

	// Wait for the goroutine to not leak
	<-slowCh

	// 9. WORKER POOL PATTERN

	fmt.Println("\n--- Worker Pool ---")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go workerPool(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for i := 1; i <= 5; i++ {
		fmt.Printf("Result: %d\n", <-results)
	}

	// 10. FAN-OUT / FAN-IN PATTERN

	fmt.Println("\n--- Fan-Out / Fan-In ---")

	input := make(chan int)
	output := fanIn(
		square(input),
		square(input),
		square(input),
	)

	// Send input
	go func() {
		for i := 1; i <= 9; i++ {
			input <- i
		}
		close(input)
	}()

	// Collect output
	for result := range output {
		fmt.Printf("%d ", result)
	}
	fmt.Println()

	// 11. PIPELINE PATTERN

	fmt.Println("\n--- Pipeline ---")

	// Stage 1: Generate numbers
	nums := gen(1, 2, 3, 4, 5)

	// Stage 2: Square them
	squared := sq(nums)

	// Stage 3: Consume
	for n := range squared {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	// 12. MUTEX VS CHANNELS

	fmt.Println("\n--- Mutex vs Channels ---")

	// Mutex approach
	counter := &SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc("key")
		}()
	}
	wg.Wait()
	fmt.Println("Mutex counter:", counter.Value("key"))

	// Channel approach
	counterCh := make(chan int)
	resultCh := make(chan int)

	go func() {
		count := 0
		for delta := range counterCh {
			count += delta
		}
		resultCh <- count
	}()

	for i := 0; i < 100; i++ {
		counterCh <- 1
	}
	close(counterCh)
	fmt.Println("Channel counter:", <-resultCh)

	// 13. QUIT CHANNEL PATTERN

	fmt.Println("\n--- Quit Channel ---")

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Worker: shutting down")
				return
			default:
				fmt.Println("Worker: working...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(150 * time.Millisecond)
	close(quit) // signal worker to stop
	time.Sleep(50 * time.Millisecond)

	fmt.Println("\n--- End of Demo ---")
}

// HELPER FUNCTIONS

// send only accepts send-only channel
func send(ch chan<- int) {
	ch <- 42
}

// receive only accepts receive-only channel
func receive(ch <-chan int) {
	fmt.Println("Received:", <-ch)
}

// Worker for worker pool
func workerPool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(50 * time.Millisecond)
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

// gen generates numbers
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq squares numbers from input channel
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// SAFE COUNTER WITH MUTEX

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}
