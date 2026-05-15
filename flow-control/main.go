package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
  FLOW CONTROL IN GO

KEY POINTS:
  - Go has only ONE loop construct: for
  - if/else requires curly braces but not parentheses
  - switch is more powerful than in C/Java (no automatic fallthrough)
  - defer delays execution until surrounding function returns

LOOPS:
  for init; condition; post { }   // C-style for
  for condition { }               // while loop
  for { }                         // infinite loop
  for i, v := range collection { } // range loop

CONDITIONALS:
  if condition { } else { }
  switch value { case x: ... }
  switch { case condition: ... }  // tagless switch

*/

func main() {
	// 1. FOR LOOP - BASIC

	fmt.Println("--- For Loop (C-style) ---")

	for i := 0; i < 5; i++ {
		fmt.Printf("i=%d ", i)
	}
	fmt.Println()

	// With just condition (like while)
	fmt.Println("\n--- While-style Loop ---")
	count := 0
	for count < 3 {
		fmt.Println("count:", count)
		count++
	}

	// Infinite loop (use break to exit)
	fmt.Println("\n--- Infinite Loop with Break ---")
	n := 0
	for {
		fmt.Println("n:", n)
		n++
		if n >= 3 {
			break
		}
	}

	// 2. RANGE LOOP

	fmt.Println("\n--- Range Loop ---")

	// Over slice
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, fruit)
	}

	// Value only (ignore index)
	fmt.Print("Values: ")
	for _, fruit := range fruits {
		fmt.Print(fruit, " ")
	}
	fmt.Println()

	// Index only
	fmt.Print("Indices: ")
	for i := range fruits {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Over map
	fmt.Println("\nRange over map:")
	scores := map[string]int{"Alice": 95, "Bob": 87}
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Over string (iterates over runes)
	fmt.Print("Runes in 'Go语言': ")
	for i, r := range "Go语言" {
		fmt.Printf("[%d:%c] ", i, r)
	}
	fmt.Println()

	// 3. BREAK AND CONTINUE

	fmt.Println("\n--- Break and Continue ---")

	// Continue skips to next iteration
	fmt.Print("Skip evens: ")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Break exits the loop
	fmt.Print("Break at 5: ")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Labeled break (for nested loops)
	fmt.Println("\nLabeled break:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking out of both loops")
				break outer
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// 4. IF/ELSE

	fmt.Println("\n--- If/Else ---")

	x := 10

	// Basic if
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// If-else
	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	// If-else if-else chain
	if x < 0 {
		fmt.Println("x is negative")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is positive")
	}

	// If with short statement
	if y := x * 2; y > 15 {
		fmt.Println("y (x*2) is greater than 15, y =", y)
	}

	// y is not accessible here (scoped to if block)

	// 5. SWITCH STATEMENT

	fmt.Println("\n--- Switch Statement ---")

	// Basic switch
	day := "Wednesday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Wednesday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday": // multiple values
		fmt.Println("Weekend!")
	default:
		fmt.Println("Regular day")
	}

	// Switch with expression
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Switch with short statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	case "windows":
		fmt.Println("Running on Windows")
	default:
		fmt.Println("Running on", os)
	}

	// Type switch
	fmt.Println("\nType switch:")
	printType(42)
	printType("hello")
	printType(3.14)
	printType([]int{1, 2, 3})

	// Fallthrough (explicit, unlike C)
	fmt.Println("\nFallthrough:")
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two (fallthrough)")
		fallthrough
	case 3:
		fmt.Println("Three (fallthrough)")
	}

	// 6. DEFER

	fmt.Println("\n--- Defer ---")

	// Defer delays execution until function returns
	fmt.Println("Counting...")
	deferDemo()

	// Practical: resource cleanup
	fmt.Println("\nPractical defer:")
	processFile()

	// Defer captures values at defer time
	fmt.Println("\nDefer value capture:")
	captureDemo()

	// 7. PRACTICAL EXAMPLES

	fmt.Println("\n--- Practical: FizzBuzz ---")
	for i := 1; i <= 15; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	fmt.Println("\n--- Practical: Find Prime ---")
	for _, n := range []int{2, 7, 10, 13, 20, 29} {
		if isPrime(n) {
			fmt.Printf("%d is prime\n", n)
		} else {
			fmt.Printf("%d is not prime\n", n)
		}
	}
}

func printType(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("  int: %d\n", t)
	case string:
		fmt.Printf("  string: %q\n", t)
	case float64:
		fmt.Printf("  float64: %.2f\n", t)
	default:
		fmt.Printf("  unknown: %T\n", t)
	}
}

func deferDemo() {
	defer fmt.Println("3 (deferred last, runs first)")
	defer fmt.Println("2 (deferred second)")
	defer fmt.Println("1 (deferred first, runs last)")
	fmt.Println("0 (runs immediately)")
}

func processFile() {
	fmt.Println("Opening file...")
	defer fmt.Println("Closing file... (cleanup)")

	fmt.Println("Reading file...")
	fmt.Println("Processing data...")
}

func captureDemo() {
	x := 10

	// Value captured at defer time
	defer fmt.Printf("  Captured value: %d\n", x)

	// Closure sees final value
	defer func() {
		fmt.Printf("  Closure sees: %d\n", x)
	}()

	x = 20
	fmt.Printf("  Current x: %d\n", x)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
