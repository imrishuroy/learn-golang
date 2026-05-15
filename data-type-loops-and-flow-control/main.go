package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
  DATA TYPES, LOOPS, AND FLOW CONTROL IN GO

KEY POINTS:
  - Go is statically typed with type inference
  - Only one loop construct: for (no while, do-while)
  - Switch doesn't need break (no fallthrough by default)
  - Defer delays execution until function returns

DATA TYPES:
  Basic: bool, string, int, float64, byte, rune
  Composite: array, slice, map, struct
  Reference: pointer, channel, function

LOOPS:
  for init; cond; post { }   // C-style for
  for condition { }          // while
  for { }                    // infinite
  for i, v := range coll { } // range

FLOW CONTROL:
  if, if-else, if-else if-else
  switch (with/without expression)
  select (for channels)
  defer (cleanup)

*/

// Constants
const (
	Pi        = 3.14159
	MaxUsers  = 100
	AppName   = "GoTutorial"
)

func main() {
	// ==================== DATA TYPES ====================

	fmt.Println("========== DATA TYPES ==========")

	// 1. Variable Declaration
	fmt.Println("\n--- Variable Declaration ---")

	var a int = 10           // Explicit type
	var b = 20               // Type inferred
	c := 30                  // Short declaration (most common)
	var d int                // Zero value (0)

	fmt.Printf("a=%d, b=%d, c=%d, d=%d\n", a, b, c, d)

	// Multiple variables
	var x, y, z int = 1, 2, 3
	p, q, r := "hello", 42, true
	fmt.Println("Multiple:", x, y, z)
	fmt.Println("Different types:", p, q, r)

	// 2. Basic Types
	fmt.Println("\n--- Basic Types ---")

	var (
		isActive bool    = true
		name     string  = "Go"
		age      int     = 10
		pi       float64 = 3.14159
		char     rune    = 'G'          // Unicode code point
		data     byte    = 255          // Alias for uint8
	)

	fmt.Printf("bool: %v\n", isActive)
	fmt.Printf("string: %s\n", name)
	fmt.Printf("int: %d\n", age)
	fmt.Printf("float64: %.2f\n", pi)
	fmt.Printf("rune: %c (value: %d)\n", char, char)
	fmt.Printf("byte: %d\n", data)

	// 3. Type Conversion
	fmt.Println("\n--- Type Conversion ---")

	intVal := 42
	floatVal := float64(intVal)
	uintVal := uint(intVal)

	fmt.Printf("int: %d → float64: %.1f → uint: %d\n", intVal, floatVal, uintVal)

	// Float to int truncates
	pi2 := 3.99
	truncated := int(pi2)
	fmt.Printf("float %.2f → int %d (truncates)\n", pi2, truncated)

	// 4. Constants
	fmt.Println("\n--- Constants ---")

	fmt.Println("Pi:", Pi)
	fmt.Println("MaxUsers:", MaxUsers)
	fmt.Println("AppName:", AppName)

	// iota - auto-incrementing
	const (
		Sunday = iota  // 0
		Monday         // 1
		Tuesday        // 2
		Wednesday      // 3
	)
	fmt.Println("Wednesday =", Wednesday)

	// iota with expression
	const (
		_  = iota             // 0 (ignored)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20
		GB                    // 1 << 30
	)
	fmt.Printf("KB=%d, MB=%d, GB=%d\n", KB, MB, GB)

	// ==================== LOOPS ====================

	fmt.Println("\n========== LOOPS ==========")

	// 5. C-style For Loop
	fmt.Println("\n--- C-style For ---")

	fmt.Print("0 to 4: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 6. While-style Loop
	fmt.Println("\n--- While-style ---")

	count := 0
	fmt.Print("While count < 3: ")
	for count < 3 {
		fmt.Print(count, " ")
		count++
	}
	fmt.Println()

	// 7. Infinite Loop with Break
	fmt.Println("\n--- Infinite Loop ---")

	n := 0
	fmt.Print("Break at 3: ")
	for {
		fmt.Print(n, " ")
		n++
		if n >= 3 {
			break
		}
	}
	fmt.Println()

	// 8. Range Loop
	fmt.Println("\n--- Range Loop ---")

	nums := []int{10, 20, 30, 40, 50}

	fmt.Print("Index and value: ")
	for i, v := range nums {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	fmt.Print("Value only: ")
	for _, v := range nums {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Index only: ")
	for i := range nums {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Range over string (iterates runes)
	fmt.Print("Runes in 'Go语言': ")
	for i, r := range "Go语言" {
		fmt.Printf("[%d:%c] ", i, r)
	}
	fmt.Println()

	// 9. Continue
	fmt.Println("\n--- Continue ---")

	fmt.Print("Skip evens: ")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 10. Labeled Break
	fmt.Println("\n--- Labeled Break ---")

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

	// ==================== FLOW CONTROL ====================

	fmt.Println("\n========== FLOW CONTROL ==========")

	// 11. If-Else
	fmt.Println("\n--- If-Else ---")

	number := 10
	if number > 0 {
		fmt.Println("Positive")
	} else if number < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	// If with short statement
	if val := number * 2; val > 15 {
		fmt.Printf("%d * 2 = %d (greater than 15)\n", number, val)
	}

	// 12. Switch
	fmt.Println("\n--- Switch ---")

	day := "Wednesday"
	switch day {
	case "Monday":
		fmt.Println("Start of week")
	case "Wednesday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday": // Multiple values
		fmt.Println("Weekend!")
	default:
		fmt.Println("Regular day")
	}

	// Switch without expression (like if-else chain)
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
	default:
		fmt.Println("Running on", os)
	}

	// Fallthrough (explicit)
	fmt.Println("\nFallthrough example:")
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two (fallthrough)")
	case 3:
		fmt.Println("Three")
	}

	// 13. Defer
	fmt.Println("\n--- Defer ---")

	fmt.Println("Counting with defer:")
	deferDemo()

	// Defer captures values at defer time
	fmt.Println("\nDefer value capture:")
	captureDemo()

	// 14. Printf Specifiers
	fmt.Println("\n--- Printf Specifiers ---")

	fmt.Printf("%%v (value): %v\n", nums)
	fmt.Printf("%%+v (with fields): %+v\n", struct{ X, Y int }{1, 2})
	fmt.Printf("%%T (type): %T\n", nums)
	fmt.Printf("%%d (decimal): %d\n", 42)
	fmt.Printf("%%f (float): %f\n", 3.14159)
	fmt.Printf("%%.2f (precision): %.2f\n", 3.14159)
	fmt.Printf("%%s (string): %s\n", "hello")
	fmt.Printf("%%q (quoted): %q\n", "hello")
	fmt.Printf("%%t (bool): %t\n", true)
	fmt.Printf("%%p (pointer): %p\n", &nums)
}

func deferDemo() {
	defer fmt.Println("3 (deferred last)")
	defer fmt.Println("2 (deferred second)")
	defer fmt.Println("1 (deferred first)")
	fmt.Println("0 (runs immediately)")
}

func captureDemo() {
	x := 10
	defer fmt.Printf("  Captured at defer time: %d\n", x)
	defer func() {
		fmt.Printf("  Closure sees final value: %d\n", x)
	}()
	x = 20
	fmt.Printf("  Current x: %d\n", x)
}
