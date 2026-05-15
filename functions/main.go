package main

import (
	"fmt"
	"strings"
)

/*
  FUNCTIONS IN GO

KEY POINTS:
  - Functions are first-class citizens (can be passed as values)
  - Multiple return values are idiomatic
  - Named return values for documentation
  - Variadic functions accept variable arguments
  - Closures capture variables from outer scope

SYNTAX:
  func name(params) returnType { }
  func name(params) (ret1, ret2) { }
  func name(params) (named Type) { }
  func name(args ...Type) { }

BEST PRACTICES:
  - Return errors as the last return value
  - Keep functions small and focused
  - Use named returns for complex functions

*/

// ChargeSession for examples
type ChargeSession struct {
	Id    string
	Watts int
	Vin   string
}

func main() {
	// 1. BASIC FUNCTIONS

	fmt.Println("--- Basic Functions ---")

	greet("Alice")
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// 2. MULTIPLE RETURN VALUES

	fmt.Println("\n--- Multiple Return Values ---")

	sum, product := sumAndProduct(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// Common pattern: value and error
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", quotient)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 3. NAMED RETURN VALUES

	fmt.Println("\n--- Named Return Values ---")

	area, perimeter := rectangleStats(10, 5)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", area, perimeter)

	// 4. VARIADIC FUNCTIONS

	fmt.Println("\n--- Variadic Functions ---")

	fmt.Println("Sum():", sumAll())
	fmt.Println("Sum(1):", sumAll(1))
	fmt.Println("Sum(1,2,3):", sumAll(1, 2, 3))
	fmt.Println("Sum(1..5):", sumAll(1, 2, 3, 4, 5))

	// Pass slice to variadic using ...
	nums := []int{10, 20, 30}
	fmt.Println("Sum(slice...):", sumAll(nums...))

	// 5. HIGHER-ORDER FUNCTIONS

	fmt.Println("\n--- Higher-Order Functions ---")

	// Function as argument
	numbers := []int{1, 2, 3, 4, 5}

	doubled := apply(numbers, func(x int) int {
		return x * 2
	})
	fmt.Println("Doubled:", doubled)

	squared := apply(numbers, func(x int) int {
		return x * x
	})
	fmt.Println("Squared:", squared)

	// 6. FUNCTIONS AS VALUES

	fmt.Println("\n--- Functions as Values ---")

	// Assign function to variable
	operation := add
	fmt.Println("operation(3, 4):", operation(3, 4))

	operation = multiply
	fmt.Println("operation(3, 4):", operation(3, 4))

	fmt.Printf("Type: %T\n", operation)

	// 7. CLOSURES

	fmt.Println("\n--- Closures ---")

	// Closure captures variable from outer scope
	counter := makeCounter()
	fmt.Println("counter():", counter()) // 1
	fmt.Println("counter():", counter()) // 2
	fmt.Println("counter():", counter()) // 3

	// Each closure has its own state
	counter2 := makeCounter()
	fmt.Println("counter2():", counter2()) // 1 (fresh)
	fmt.Println("counter():", counter())   // 4 (continues)

	// 8. FUNCTION FACTORY

	fmt.Println("\n--- Function Factory ---")

	addFive := makeAdder(5)
	addTen := makeAdder(10)

	fmt.Println("addFive(3):", addFive(3))   // 8
	fmt.Println("addTen(3):", addTen(3))     // 13

	// 9. CLOSURE GOTCHA - LOOP VARIABLES

	fmt.Println("\n--- Closure Gotcha ---")

	// WRONG: All closures share same loop variable
	var wrongFuncs []func()
	for i := 0; i < 3; i++ {
		wrongFuncs = append(wrongFuncs, func() {
			fmt.Printf("wrong: i=%d ", i)
		})
	}
	fmt.Print("Wrong way: ")
	for _, fn := range wrongFuncs {
		fn() // Prints: 3 3 3
	}
	fmt.Println()

	// CORRECT: Capture loop variable
	var rightFuncs []func()
	for i := 0; i < 3; i++ {
		i := i // Create new variable in each iteration
		rightFuncs = append(rightFuncs, func() {
			fmt.Printf("right: i=%d ", i)
		})
	}
	fmt.Print("Right way: ")
	for _, fn := range rightFuncs {
		fn() // Prints: 0 1 2
	}
	fmt.Println()

	// 10. FILTER FUNCTION

	fmt.Println("\n--- Filter Function ---")

	sessions := []ChargeSession{
		{Id: "CS001", Watts: 420, Vin: "ABC123"},
		{Id: "CS002", Watts: 350, Vin: "XYZ789"},
		{Id: "CS003", Watts: 500, Vin: "DEF456"},
	}

	// Filter sessions with Watts > 400
	highPower := filter(sessions, func(cs ChargeSession) bool {
		return cs.Watts > 400
	})

	fmt.Println("High power sessions (>400W):")
	for _, s := range highPower {
		fmt.Printf("  %s: %dW\n", s.Id, s.Watts)
	}

	// 11. DECORATOR PATTERN

	fmt.Println("\n--- Decorator Pattern ---")

	hello := func(name string) string {
		return "Hello, " + name
	}

	loudHello := withUppercase(hello)
	excitedHello := withExclamation(hello)

	fmt.Println(hello("World"))
	fmt.Println(loudHello("World"))
	fmt.Println(excitedHello("World"))

	// 12. IMMEDIATELY INVOKED FUNCTION

	fmt.Println("\n--- Immediately Invoked ---")

	result2 := func(a, b int) int {
		return a + b
	}(10, 20) // Call immediately

	fmt.Println("10 + 20 =", result2)
}

// BASIC FUNCTIONS

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// MULTIPLE RETURNS

func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// NAMED RETURNS

func rectangleStats(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // Naked return
}

// VARIADIC

func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// HIGHER-ORDER

func apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

// CLOSURES

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func makeAdder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

// FILTER

func filter(sessions []ChargeSession, predicate func(ChargeSession) bool) []ChargeSession {
	var result []ChargeSession
	for _, s := range sessions {
		if predicate(s) {
			result = append(result, s)
		}
	}
	return result
}

// DECORATORS

func withUppercase(fn func(string) string) func(string) string {
	return func(s string) string {
		return strings.ToUpper(fn(s))
	}
}

func withExclamation(fn func(string) string) func(string) string {
	return func(s string) string {
		return fn(s) + "!"
	}
}
