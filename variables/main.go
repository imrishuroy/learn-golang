package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
  VARIABLES AND DATA TYPES IN GO

KEY POINTS:
  - Go is statically typed: variable types are fixed at compile time
  - Type inference: Go can often deduce the type from the value
  - Zero values: uninitialized variables get their type's zero value
  - Constants are immutable values known at compile time

DECLARATION SYNTAX:
  var name Type              // explicit type, zero value
  var name Type = value      // explicit type with value
  var name = value           // type inference
  name := value              // short declaration (most common)

BASIC TYPES:
  bool                       // true or false
  string                     // UTF-8 text
  int, int8, int16, int32, int64
  uint, uint8, uint16, uint32, uint64
  float32, float64
  complex64, complex128
  byte (alias for uint8)
  rune (alias for int32, represents Unicode code point)

*/

func main() {
	// 1. VARIABLE DECLARATIONS

	fmt.Println("--- Variable Declarations ---")

	// Method 1: var with explicit type (gets zero value)
	var age int
	var name string
	var isActive bool
	fmt.Printf("Zero values: age=%d, name=%q, isActive=%v\n", age, name, isActive)

	// Method 2: var with explicit type and value
	var score int = 100
	fmt.Println("Explicit type:", score)

	// Method 3: var with type inference
	var city = "New York"
	fmt.Printf("Type inferred: %s (type: %T)\n", city, city)

	// Method 4: Short declaration := (most common)
	country := "USA"
	population := 331000000
	fmt.Printf("Short declaration: %s, pop: %d\n", country, population)

	// Multiple variables at once
	var x, y, z int = 1, 2, 3
	a, b, c := "hello", 42, true
	fmt.Println("Multiple vars:", x, y, z)
	fmt.Println("Multiple short:", a, b, c)

	// Block declaration
	var (
		firstName = "John"
		lastName  = "Doe"
		age2      = 30
	)
	fmt.Println("Block declaration:", firstName, lastName, age2)

	// 2. NUMERIC TYPES

	fmt.Println("\n--- Numeric Types ---")

	// Integers
	var i8 int8 = 127                  // -128 to 127
	var i16 int16 = 32767              // -32768 to 32767
	var i32 int32 = 2147483647         // ~2 billion
	var i64 int64 = 9223372036854775807 // ~9 quintillion
	var i int = 42                     // platform dependent (32 or 64 bit)

	fmt.Printf("int8=%d, int16=%d, int32=%d\n", i8, i16, i32)
	fmt.Printf("int64=%d, int=%d\n", i64, i)

	// Unsigned integers
	var u8 uint8 = 255     // 0 to 255
	var u16 uint16 = 65535 // 0 to 65535
	fmt.Printf("uint8=%d, uint16=%d\n", u8, u16)

	// Floating point
	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793
	fmt.Printf("float32=%.5f, float64=%.15f\n", f32, f64)

	// Complex numbers
	var c64 complex64 = 1 + 2i
	var c128 complex128 = complex(3, 4)
	fmt.Printf("complex64=%v, complex128=%v\n", c64, c128)
	fmt.Printf("real=%f, imag=%f\n", real(c128), imag(c128))

	// 3. STRINGS AND RUNES

	fmt.Println("\n--- Strings and Runes ---")

	// Strings are immutable sequences of bytes
	str := "Hello, World!"
	fmt.Println("String:", str)
	fmt.Println("Length (bytes):", len(str))
	fmt.Println("First byte:", str[0]) // 72 (ASCII for 'H')

	// Rune = Unicode code point (int32)
	r := '世'
	fmt.Printf("Rune: %c, value: %d, type: %T\n", r, r, r)

	// String with Unicode
	greeting := "Hello, 世界!"
	fmt.Println("Unicode string:", greeting)
	fmt.Println("Byte length:", len(greeting))

	// Iterate over runes (not bytes)
	fmt.Print("Runes: ")
	for _, char := range greeting {
		fmt.Printf("%c ", char)
	}
	fmt.Println()

	// Raw strings (no escape sequences)
	raw := `This is a raw string.
It can span multiple lines.
\n is literal, not a newline.`
	fmt.Println("Raw string:", raw)

	// 4. BOOLEANS

	fmt.Println("\n--- Booleans ---")

	var isTrue bool = true
	var isFalse bool = false
	fmt.Println("true:", isTrue, "false:", isFalse)

	// Boolean operations
	fmt.Println("AND:", true && false)  // false
	fmt.Println("OR:", true || false)   // true
	fmt.Println("NOT:", !true)          // false

	// Comparison operators return bool
	fmt.Println("5 > 3:", 5 > 3)        // true
	fmt.Println("5 == 5:", 5 == 5)      // true
	fmt.Println("5 != 5:", 5 != 5)      // false

	// 5. ZERO VALUES

	fmt.Println("\n--- Zero Values ---")

	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	var zeroPointer *int

	fmt.Printf("int: %d\n", zeroInt)
	fmt.Printf("float64: %f\n", zeroFloat)
	fmt.Printf("string: %q\n", zeroString)
	fmt.Printf("bool: %v\n", zeroBool)
	fmt.Printf("pointer: %v\n", zeroPointer)

	// 6. CONSTANTS

	fmt.Println("\n--- Constants ---")

	// Constants are immutable
	const Pi = 3.14159
	const MaxUsers = 100
	const Greeting = "Hello"

	fmt.Println("Pi:", Pi)
	fmt.Println("MaxUsers:", MaxUsers)
	fmt.Println("Greeting:", Greeting)

	// Typed vs untyped constants
	const untypedInt = 42      // untyped - can be used with any numeric type
	const typedInt int = 42    // typed - only int

	var f float64 = untypedInt // works: untyped adapts
	// var f2 float64 = typedInt // error: cannot use int as float64
	_ = f

	// Constant expressions
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)
	fmt.Println("1 GB =", GB, "bytes")

	// iota - auto-incrementing constant generator
	const (
		Sunday = iota  // 0
		Monday         // 1
		Tuesday        // 2
		Wednesday      // 3
		Thursday       // 4
		Friday         // 5
		Saturday       // 6
	)
	fmt.Println("Wednesday =", Wednesday)

	// iota patterns
	const (
		_  = iota             // 0 (ignored)
		KB2 = 1 << (10 * iota) // 1 << 10 = 1024
		MB2                   // 1 << 20 = 1048576
		GB2                   // 1 << 30
	)
	fmt.Println("KB2:", KB2, "MB2:", MB2)

	// 7. TYPE CONVERSIONS

	fmt.Println("\n--- Type Conversions ---")

	// Go requires explicit conversions
	var intVal int = 42
	var floatVal float64 = float64(intVal)
	var uintVal uint = uint(intVal)

	fmt.Printf("int=%d, float64=%.1f, uint=%d\n", intVal, floatVal, uintVal)

	// Float to int (truncates)
	var pi float64 = 3.99
	var truncated int = int(pi)
	fmt.Printf("%.2f truncated = %d\n", pi, truncated)

	// String conversions
	var num int = 65
	var char string = string(num) // converts to Unicode character, not "65"
	fmt.Printf("int %d to string: %q (Unicode char)\n", num, char)

	// For number to string, use strconv or fmt.Sprintf
	numStr := fmt.Sprintf("%d", 42)
	fmt.Println("Number to string:", numStr)

	// 8. TYPE INSPECTION

	fmt.Println("\n--- Type Inspection ---")

	values := []interface{}{42, 3.14, "hello", true, nil}

	for _, v := range values {
		fmt.Printf("Value: %-10v Type: %T\n", v, v)
	}

	// Using reflect package
	var myVar float64 = 3.14159
	t := reflect.TypeOf(myVar)
	fmt.Println("reflect.TypeOf:", t.Name(), "Kind:", t.Kind())

	// 9. SPECIAL NUMERIC OPERATIONS

	fmt.Println("\n--- Special Numeric Operations ---")

	// Integer division vs float division
	fmt.Println("7 / 3 (int):", 7/3)           // 2
	fmt.Println("7.0 / 3.0 (float):", 7.0/3.0) // 2.333...

	// Modulo
	fmt.Println("7 % 3:", 7%3) // 1

	// Increment/Decrement (statements, not expressions)
	counter := 0
	counter++
	counter++
	counter--
	fmt.Println("Counter after ++, ++, --:", counter)

	// Math package for advanced operations
	fmt.Println("math.Sqrt(16):", math.Sqrt(16))
	fmt.Println("math.Pow(2, 10):", math.Pow(2, 10))
	fmt.Println("math.Max(5, 3):", math.Max(5, 3))
	fmt.Println("math.Abs(-5):", math.Abs(-5))

	// 10. PRACTICAL: TEMPERATURE CONVERSION

	fmt.Println("\n--- Practical: Temperature Conversion ---")

	celsius := 100.0
	fahrenheit := celsius*9/5 + 32
	kelvin := celsius + 273.15

	fmt.Printf("%.1f°C = %.1f°F = %.2fK\n", celsius, fahrenheit, kelvin)
}
