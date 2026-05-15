package main

import (
	"fmt"
	"sync"
	"time"
)

/*
  MUTEX AND SYNCHRONIZATION IN GO

KEY POINTS:
  - Mutex = Mutual Exclusion - prevents concurrent access to shared data
  - sync.Mutex: Only one goroutine can hold the lock
  - sync.RWMutex: Multiple readers OR one writer
  - sync.Once: Execute code exactly once
  - sync.Cond: Condition variables for signaling

WHEN TO USE:
  - Mutex: Protecting shared state from data races
  - RWMutex: Many reads, few writes (read-heavy workloads)
  - Channels: Communication between goroutines
  - Atomic: Simple counters and flags

RULES:
  - Always unlock in same function (use defer)
  - Don't copy mutex after first use
  - Avoid holding locks while doing I/O
  - Keep critical sections short

*/

func main() {
	// 1. DATA RACE PROBLEM

	fmt.Println("--- Data Race Problem ---")

	// Without mutex: data race!
	counter1 := 0
	var wg1 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			counter1++ // NOT thread-safe!
		}()
	}
	wg1.Wait()
	fmt.Printf("Counter without mutex: %d (expected 1000, likely wrong)\n", counter1)

	// 2. MUTEX SOLUTION

	fmt.Println("\n--- Mutex Solution ---")

	counter2 := 0
	var mu sync.Mutex
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter2++ // Protected by mutex
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("Counter with mutex: %d (always 1000)\n", counter2)

	// 3. MUTEX WITH DEFER

	fmt.Println("\n--- Mutex with Defer ---")

	var safeMu sync.Mutex
	safeCounter := 0

	increment := func() {
		safeMu.Lock()
		defer safeMu.Unlock() // Guaranteed unlock even if panic

		safeCounter++
		// Any panic here would still unlock
	}

	var wg3 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			increment()
		}()
	}
	wg3.Wait()
	fmt.Println("Safe counter:", safeCounter)

	// 4. SAFE COUNTER STRUCT

	fmt.Println("\n--- Safe Counter Struct ---")

	sc := &SafeCounter{}

	var wg4 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			sc.Inc()
		}()
	}
	wg4.Wait()
	fmt.Println("SafeCounter value:", sc.Value())

	// 5. RWMUTEX - READ/WRITE MUTEX

	fmt.Println("\n--- RWMutex ---")

	cache := NewCache()

	// Multiple writers
	var wg5 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg5.Add(1)
		go func(n int) {
			defer wg5.Done()
			key := fmt.Sprintf("key%d", n)
			cache.Set(key, n*10)
			fmt.Printf("Set %s = %d\n", key, n*10)
		}(i)
	}
	wg5.Wait()

	// Multiple concurrent readers
	fmt.Println("\nConcurrent reads:")
	for i := 0; i < 5; i++ {
		wg5.Add(1)
		go func(n int) {
			defer wg5.Done()
			key := fmt.Sprintf("key%d", n)
			val := cache.Get(key)
			fmt.Printf("Get %s = %d\n", key, val)
		}(i)
	}
	wg5.Wait()

	// 6. SYNC.ONCE

	fmt.Println("\n--- sync.Once ---")

	var once sync.Once
	var wg6 sync.WaitGroup

	initFunc := func() {
		fmt.Println("Initialization runs only once!")
	}

	// Try to initialize from multiple goroutines
	for i := 0; i < 5; i++ {
		wg6.Add(1)
		go func(n int) {
			defer wg6.Done()
			fmt.Printf("Goroutine %d calling once.Do\n", n)
			once.Do(initFunc) // Only first call executes
		}(i)
	}
	wg6.Wait()

	// 7. SYNC.ONCE PRACTICAL: SINGLETON

	fmt.Println("\n--- Singleton Pattern ---")

	var wg7 sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg7.Add(1)
		go func(n int) {
			defer wg7.Done()
			db := GetDB()
			fmt.Printf("Goroutine %d got DB: %s\n", n, db.name)
		}(i)
	}
	wg7.Wait()

	// 8. SYNC.MAP

	fmt.Println("\n--- sync.Map ---")

	var sm sync.Map

	// Store values concurrently
	var wg8 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg8.Add(1)
		go func(n int) {
			defer wg8.Done()
			sm.Store(fmt.Sprintf("key%d", n), n*100)
		}(i)
	}
	wg8.Wait()

	// Load values
	fmt.Println("sync.Map contents:")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("  %s: %v\n", key, value)
		return true // continue iteration
	})

	// Load with existence check
	if val, ok := sm.Load("key2"); ok {
		fmt.Println("Found key2:", val)
	}

	// LoadOrStore
	actual, loaded := sm.LoadOrStore("key10", 999)
	fmt.Printf("LoadOrStore key10: value=%v, wasLoaded=%v\n", actual, loaded)

	// 9. SYNC.POOL

	fmt.Println("\n--- sync.Pool ---")

	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new buffer")
			return make([]byte, 1024)
		},
	}

	// First get creates new object
	buf1 := pool.Get().([]byte)
	fmt.Printf("Got buffer of size %d\n", len(buf1))

	// Put it back
	pool.Put(buf1)

	// Second get reuses existing
	buf2 := pool.Get().([]byte)
	fmt.Printf("Got buffer of size %d (reused)\n", len(buf2))

	// 10. PRACTICAL: CONCURRENT MAP ACCESS

	fmt.Println("\n--- Practical: User Session Manager ---")

	sessions := NewSessionManager()

	// Concurrent session creation
	var wg10 sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg10.Add(1)
		go func(userID int) {
			defer wg10.Done()
			token := sessions.CreateSession(userID)
			fmt.Printf("User %d session: %s\n", userID, token)
		}(i + 1)
	}
	wg10.Wait()

	fmt.Println("\nActive sessions:", sessions.Count())

	// 11. DEADLOCK EXAMPLE

	fmt.Println("\n--- Deadlock Prevention ---")

	fmt.Println("Common deadlock causes:")
	fmt.Println("  1. Lock ordering inconsistency")
	fmt.Println("  2. Forgetting to unlock")
	fmt.Println("  3. Recursive locking (Mutex is not reentrant)")
	fmt.Println("  4. Goroutine waiting on itself")

	// Demonstrate proper lock ordering
	demonstrateLockOrdering()
}

// SAFE COUNTER STRUCT

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// CACHE WITH RWMUTEX

type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]int)}
}

func (c *Cache) Get(key string) int {
	c.mu.RLock() // Read lock - multiple readers allowed
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock() // Write lock - exclusive access
	defer c.mu.Unlock()
	c.data[key] = value
}

// SINGLETON PATTERN

type Database struct {
	name string
}

var (
	dbInstance *Database
	dbOnce     sync.Once
)

func GetDB() *Database {
	dbOnce.Do(func() {
		fmt.Println("Initializing database connection...")
		time.Sleep(10 * time.Millisecond) // Simulate slow init
		dbInstance = &Database{name: "MainDB"}
	})
	return dbInstance
}

// SESSION MANAGER

type SessionManager struct {
	mu       sync.RWMutex
	sessions map[int]string
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[int]string),
	}
}

func (sm *SessionManager) CreateSession(userID int) string {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	token := fmt.Sprintf("session_%d_%d", userID, time.Now().UnixNano())
	sm.sessions[userID] = token
	return token
}

func (sm *SessionManager) Count() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.sessions)
}

// LOCK ORDERING

type Account struct {
	mu      sync.Mutex
	id      int
	balance int
}

func demonstrateLockOrdering() {
	fmt.Println("\nProper lock ordering for transfer:")

	a1 := &Account{id: 1, balance: 100}
	a2 := &Account{id: 2, balance: 50}

	transfer := func(from, to *Account, amount int) {
		// Always lock in consistent order (by ID) to prevent deadlock
		first, second := from, to
		if from.id > to.id {
			first, second = to, from
		}

		first.mu.Lock()
		defer first.mu.Unlock()
		second.mu.Lock()
		defer second.mu.Unlock()

		if from.balance >= amount {
			from.balance -= amount
			to.balance += amount
			fmt.Printf("Transferred %d from Account %d to Account %d\n",
				amount, from.id, to.id)
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		transfer(a1, a2, 30)
	}()

	go func() {
		defer wg.Done()
		transfer(a2, a1, 20)
	}()

	wg.Wait()
	fmt.Printf("Final balances: Account1=%d, Account2=%d\n", a1.balance, a2.balance)
}
