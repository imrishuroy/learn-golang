package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
  ANONYMOUS FUNCTIONS & CLOSURES IN GO

KEY POINTS:
  - Functions are FIRST-CLASS citizens in Go
  - Anonymous functions have no name
  - Closures capture variables from their enclosing scope
  - Functions can be assigned to variables, passed as args, returned

SYNTAX:
  func(params) returnType { body }     // anonymous function
  fn := func(x int) int { return x*2 } // assigned to variable
  func(x int) { ... }(arg)             // immediately invoked

USE CASES:
  - Callbacks and handlers
  - Goroutines: go func() { ... }()
  - Deferred execution: defer func() { ... }()
  - Custom sorting
  - Middleware patterns

*/

func main() {
	// 1. BASIC ANONYMOUS FUNCTION

	fmt.Println("--- Basic Anonymous Function ---")

	// Assign to variable
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}

	greet("Alice")
	greet("Bob")

	// Check the type
	fmt.Printf("Type of greet: %T\n", greet)

	// 2. IMMEDIATELY INVOKED FUNCTION (IIFE)

	fmt.Println("\n--- Immediately Invoked ---")

	// Execute immediately with ()
	result := func(a, b int) int {
		return a + b
	}(10, 20)

	fmt.Println("10 + 20 =", result)

	// Common pattern for initialization
	func() {
		fmt.Println("This runs immediately!")
	}()

	// 3. FUNCTIONS AS PARAMETERS

	fmt.Println("\n--- Functions as Parameters ---")

	// Pass function as argument
	numbers := []int{1, 2, 3, 4, 5}

	doubled := apply(numbers, func(x int) int {
		return x * 2
	})
	fmt.Println("Doubled:", doubled)

	squared := apply(numbers, func(x int) int {
		return x * x
	})
	fmt.Println("Squared:", squared)

	// 4. FUNCTIONS RETURNING FUNCTIONS

	fmt.Println("\n--- Functions Returning Functions ---")

	// Factory pattern
	addFive := makeAdder(5)
	addTen := makeAdder(10)

	fmt.Println("addFive(3) =", addFive(3))   // 8
	fmt.Println("addTen(3) =", addTen(3))     // 13
	fmt.Println("addFive(10) =", addFive(10)) // 15

	// 5. CLOSURES - CAPTURING VARIABLES

	fmt.Println("\n--- Closures ---")

	// Closure captures variable from outer scope
	counter := makeCounter()

	fmt.Println("counter():", counter()) // 1
	fmt.Println("counter():", counter()) // 2
	fmt.Println("counter():", counter()) // 3

	// Each closure has its own captured variable
	counter2 := makeCounter()
	fmt.Println("counter2():", counter2()) // 1 (fresh counter)
	fmt.Println("counter():", counter())   // 4 (continues)

	// 6. CLOSURE GOTCHA - LOOP VARIABLES

	fmt.Println("\n--- Closure Gotcha ---")

	// WRONG: All closures share the same loop variable
	var wrongFuncs []func()
	for i := 0; i < 3; i++ {
		wrongFuncs = append(wrongFuncs, func() {
			fmt.Printf("wrong: i=%d ", i)
		})
	}
	fmt.Print("Wrong way: ")
	for _, fn := range wrongFuncs {
		fn() // prints: 3 3 3 (all see final value)
	}
	fmt.Println()

	// CORRECT: Capture loop variable as parameter
	var rightFuncs []func()
	for i := 0; i < 3; i++ {
		i := i // create new variable in each iteration
		rightFuncs = append(rightFuncs, func() {
			fmt.Printf("right: i=%d ", i)
		})
	}
	fmt.Print("Right way: ")
	for _, fn := range rightFuncs {
		fn() // prints: 0 1 2
	}
	fmt.Println()

	// 7. FUNCTION TYPES

	fmt.Println("\n--- Function Types ---")

	// Define a function type
	type MathOp func(int, int) int

	add := MathOp(func(a, b int) int { return a + b })
	sub := MathOp(func(a, b int) int { return a - b })
	mul := MathOp(func(a, b int) int { return a * b })

	fmt.Println("add(5, 3) =", add(5, 3))
	fmt.Println("sub(5, 3) =", sub(5, 3))
	fmt.Println("mul(5, 3) =", mul(5, 3))

	// Using with a calculator function
	fmt.Println("calculate(10, 5, add) =", calculate(10, 5, add))
	fmt.Println("calculate(10, 5, mul) =", calculate(10, 5, mul))

	// 8. PRACTICAL: CUSTOM SORTING

	fmt.Println("\n--- Custom Sorting ---")

	people := []struct {
		Name string
		Age  int
	}{
		{"Charlie", 35},
		{"Alice", 30},
		{"Bob", 25},
	}

	// Sort by age using anonymous function
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by age:", people)

	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("Sorted by name:", people)

	// 9. PRACTICAL: MIDDLEWARE/DECORATOR PATTERN

	fmt.Println("\n--- Decorator Pattern ---")

	// Simple greeting function
	hello := func(name string) string {
		return "Hello, " + name
	}

	// Wrap with decorators
	loudHello := withUppercase(hello)
	excitedHello := withExclamation(hello)
	loudExcitedHello := withUppercase(withExclamation(hello))

	fmt.Println(hello("World"))
	fmt.Println(loudHello("World"))
	fmt.Println(excitedHello("World"))
	fmt.Println(loudExcitedHello("World"))

	// 10. PRACTICAL: MEMOIZATION

	fmt.Println("\n--- Memoization ---")

	// Expensive calculation cached
	fib := memoize(func(n int) int {
		if n <= 1 {
			return n
		}
		a, b := 0, 1
		for i := 2; i <= n; i++ {
			a, b = b, a+b
		}
		return b
	})

	fmt.Println("fib(10):", fib(10))
	fmt.Println("fib(10) (cached):", fib(10))
	fmt.Println("fib(20):", fib(20))

	// 11. DEFER WITH ANONYMOUS FUNCTIONS

	fmt.Println("\n--- Defer with Anonymous Functions ---")

	deferExample()
}

// HELPER FUNCTIONS

// apply applies a function to each element
func apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

// makeAdder returns a function that adds n
func makeAdder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

// makeCounter returns a function that counts calls
func makeCounter() func() int {
	count := 0 // captured by the closure
	return func() int {
		count++
		return count
	}
}

// calculate uses a function type parameter
func calculate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// withUppercase is a decorator that uppercases the result
func withUppercase(fn func(string) string) func(string) string {
	return func(s string) string {
		return strings.ToUpper(fn(s))
	}
}

// withExclamation is a decorator that adds exclamation
func withExclamation(fn func(string) string) func(string) string {
	return func(s string) string {
		return fn(s) + "!"
	}
}

// memoize caches function results
func memoize(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if val, ok := cache[n]; ok {
			fmt.Printf("  (cached: %d) ", n)
			return val
		}
		result := fn(n)
		cache[n] = result
		return result
	}
}

// deferExample shows defer with anonymous function
func deferExample() {
	x := 10

	// Captures current value of x
	defer func(val int) {
		fmt.Println("Deferred with param:", val) // 10
	}(x)

	// Captures x by reference (sees final value)
	defer func() {
		fmt.Println("Deferred closure:", x) // 20
	}()

	x = 20
	fmt.Println("Current x:", x)
}
