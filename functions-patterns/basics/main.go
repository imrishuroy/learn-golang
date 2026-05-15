package main

import (
	"fmt"
	"math"
)

/*
  FUNCTIONS IN GO

KEY POINTS:
  - Functions are first-class citizens (can be passed as values)
  - Go supports multiple return values
  - Named return values act as documentation and allow naked returns
  - Variadic functions accept variable number of arguments

SYNTAX:
  func name(params) returnType { }
  func name(params) (returnType1, returnType2) { }
  func name(params) (namedReturn Type) { }
  func name(args ...Type) { }  // variadic

BEST PRACTICES:
  - Keep functions short and focused
  - Use named returns for complex functions
  - Return errors as the last return value
  - Use variadic for optional parameters

*/

func main() {
	// 1. BASIC FUNCTIONS

	fmt.Println("--- Basic Functions ---")

	// Call a simple function
	greet("Alice")
	greet("Bob")

	// Function with return value
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Function call in expression
	fmt.Println("10 * 2 =", multiply(10, 2))

	// 2. MULTIPLE PARAMETERS

	fmt.Println("\n--- Multiple Parameters ---")

	// Parameters of same type can share type
	fmt.Println("Sum:", addThree(1, 2, 3))

	// Different parameter types
	describe("Alice", 30)

	// 3. MULTIPLE RETURN VALUES

	fmt.Println("\n--- Multiple Return Values ---")

	// Two return values
	sum, product := sumAndProduct(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// Swap values
	a, b := 10, 20
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)

	// Min and max together
	min, max := minMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// 4. NAMED RETURN VALUES

	fmt.Println("\n--- Named Return Values ---")

	// Named returns act as local variables
	area, perimeter := rectangleStats(10, 5)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", area, perimeter)

	// Division with named returns
	quot, rem := divide(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", quot, rem)

	// 5. VARIADIC FUNCTIONS

	fmt.Println("\n--- Variadic Functions ---")

	// Accept any number of arguments
	fmt.Println("Sum():", sumAll())
	fmt.Println("Sum(1):", sumAll(1))
	fmt.Println("Sum(1,2,3):", sumAll(1, 2, 3))
	fmt.Println("Sum(1..10):", sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// Pass slice to variadic using ...
	nums := []int{10, 20, 30, 40}
	fmt.Println("Sum(slice...):", sumAll(nums...))

	// Printf is variadic
	fmt.Printf("Name: %s, Age: %d, Score: %.1f\n", "Bob", 25, 95.5)

	// 6. FUNCTIONS AS VALUES

	fmt.Println("\n--- Functions as Values ---")

	// Assign function to variable
	operation := add
	fmt.Println("operation(3, 4):", operation(3, 4))

	// Change which function it references
	operation = multiply
	fmt.Println("operation(3, 4):", operation(3, 4))

	// Function type
	fmt.Printf("Type: %T\n", operation)

	// 7. FUNCTIONS AS PARAMETERS

	fmt.Println("\n--- Functions as Parameters ---")

	// Pass function as argument
	fmt.Println("Apply add:", applyOperation(5, 3, add))
	fmt.Println("Apply multiply:", applyOperation(5, 3, multiply))

	// With anonymous function
	fmt.Println("Apply subtract:", applyOperation(5, 3, func(a, b int) int {
		return a - b
	}))

	// 8. FUNCTIONS RETURNING FUNCTIONS

	fmt.Println("\n--- Functions Returning Functions ---")

	// Function factory
	addFive := makeAdder(5)
	addTen := makeAdder(10)

	fmt.Println("addFive(3):", addFive(3))
	fmt.Println("addTen(3):", addTen(3))

	// Multiplier factory
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Println("double(5):", double(5))
	fmt.Println("triple(5):", triple(5))

	// 9. RECURSIVE FUNCTIONS

	fmt.Println("\n--- Recursive Functions ---")

	// Factorial
	fmt.Println("5! =", factorial(5))
	fmt.Println("10! =", factorial(10))

	// Fibonacci
	fmt.Print("Fibonacci: ")
	for i := 0; i <= 10; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()

	// 10. PRACTICAL EXAMPLES

	fmt.Println("\n--- Practical: Math Functions ---")

	// Distance between two points
	dist := distance(0, 0, 3, 4)
	fmt.Printf("Distance from (0,0) to (3,4): %.2f\n", dist)

	// Circle calculations
	radius := 5.0
	circleArea, circumference := circleStats(radius)
	fmt.Printf("Circle (r=%.1f): Area=%.2f, Circumference=%.2f\n",
		radius, circleArea, circumference)

	// Temperature conversion
	celsius := 100.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)

	// 11. BLANK IDENTIFIER

	fmt.Println("\n--- Blank Identifier ---")

	// Ignore unwanted return values
	_, onlyMax := minMax([]int{1, 2, 3, 4, 5})
	fmt.Println("Only max:", onlyMax)

	onlyMin, _ := minMax([]int{1, 2, 3, 4, 5})
	fmt.Println("Only min:", onlyMin)
}

// SIMPLE FUNCTIONS

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func addThree(a, b, c int) int {
	return a + b + c
}

func describe(name string, age int) {
	fmt.Printf("%s is %d years old\n", name, age)
}

// MULTIPLE RETURN VALUES

func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func swap(a, b int) (int, int) {
	return b, a
}

func minMax(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// NAMED RETURN VALUES

func rectangleStats(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return
}

func divide(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

// VARIADIC FUNCTION

func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// HIGHER-ORDER FUNCTIONS

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func makeMultiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

// RECURSIVE FUNCTIONS

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// PRACTICAL FUNCTIONS

func distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

func circleStats(radius float64) (area, circumference float64) {
	area = math.Pi * radius * radius
	circumference = 2 * math.Pi * radius
	return
}

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}
