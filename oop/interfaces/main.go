package main

import (
	"fmt"
	"math"
)

/*
  INTERFACES IN GO

KEY POINTS:
  - An interface defines a set of method signatures
  - A type IMPLICITLY implements an interface (no "implements" keyword)
  - Interfaces enable polymorphism and decoupling
  - Empty interface (interface{} or any) can hold any type

SYNTAX:
  type InterfaceName interface {
      MethodName(params) returnType
  }

BEST PRACTICES:
  - Keep interfaces small (1-3 methods)
  - Define interfaces where they're USED, not implemented
  - "Accept interfaces, return concrete types"
  - Use interface{}/any sparingly

COMMON INTERFACES:
  - fmt.Stringer: String() string
  - error: Error() string
  - io.Reader: Read([]byte) (int, error)
  - io.Writer: Write([]byte) (int, error)

*/

// INTERFACE DEFINITIONS

// Shape is an interface that any shape must implement
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Describer adds description capability
type Describer interface {
	Describe() string
}

// ShapeDescriber combines both interfaces
type ShapeDescriber interface {
	Shape
	Describer
}

// TYPE DEFINITIONS

// Rectangle implements Shape
type Rectangle struct {
	Width, Height float64
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

// Triangle implements Shape
type Triangle struct {
	A, B, C float64 // sides
}

func main() {
	// 1. BASIC INTERFACE USAGE

	fmt.Println("--- Basic Interface ---")

	// Both Rectangle and Circle implement Shape
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	// Can use interface type
	var s Shape

	s = rect
	fmt.Printf("Rectangle: Area=%.2f, Perimeter=%.2f\n", s.Area(), s.Perimeter())

	s = circle
	fmt.Printf("Circle: Area=%.2f, Perimeter=%.2f\n", s.Area(), s.Perimeter())

	// 2. INTERFACES ENABLE POLYMORPHISM

	fmt.Println("\n--- Polymorphism ---")

	shapes := []Shape{
		Rectangle{10, 5},
		Circle{7},
		Triangle{3, 4, 5},
	}

	for _, shape := range shapes {
		printShapeInfo(shape)
	}

	// Calculate total area
	fmt.Printf("Total area: %.2f\n", totalArea(shapes))

	// 3. IMPLICIT IMPLEMENTATION

	fmt.Println("\n--- Implicit Implementation ---")

	// No "implements" keyword needed
	// If a type has all the methods, it implements the interface
	// Square type is defined at package level with Area() and Perimeter() methods

	sq := Square{Side: 5}
	printShapeInfo(sq) // works because Square has Area() and Perimeter()

	// 4. EMPTY INTERFACE (any)

	fmt.Println("\n--- Empty Interface (any) ---")

	// interface{} or 'any' can hold any type
	var anything interface{}

	anything = 42
	fmt.Printf("int: %v (type: %T)\n", anything, anything)

	anything = "hello"
	fmt.Printf("string: %v (type: %T)\n", anything, anything)

	anything = Rectangle{10, 5}
	fmt.Printf("struct: %v (type: %T)\n", anything, anything)

	// Useful for generic containers
	mixedSlice := []any{1, "two", 3.0, true}
	fmt.Println("Mixed slice:", mixedSlice)

	// 5. TYPE ASSERTIONS

	fmt.Println("\n--- Type Assertions ---")

	var i interface{} = "hello"

	// Type assertion with ok check (safe)
	str, ok := i.(string)
	fmt.Printf("String assertion: value=%q, ok=%v\n", str, ok)

	num, ok := i.(int)
	fmt.Printf("Int assertion: value=%d, ok=%v\n", num, ok)

	// Direct assertion (panics if wrong type!)
	str2 := i.(string)
	fmt.Println("Direct assertion:", str2)

	// This would panic:
	// num2 := i.(int)

	// 6. TYPE SWITCH

	fmt.Println("\n--- Type Switch ---")

	printType(42)
	printType("hello")
	printType(3.14)
	printType(true)
	printType([]int{1, 2, 3})

	// 7. INTERFACE VALUES

	fmt.Println("\n--- Interface Values ---")

	// Interface value = (type, value) pair
	var shape Shape

	fmt.Printf("nil interface: value=%v, type=%T\n", shape, shape)

	shape = Rectangle{10, 5}
	fmt.Printf("Rectangle: value=%v, type=%T\n", shape, shape)

	shape = Circle{7}
	fmt.Printf("Circle: value=%v, type=%T\n", shape, shape)

	// 8. NIL INTERFACE VS NIL CONCRETE

	fmt.Println("\n--- Nil Interface vs Nil Concrete ---")

	var s1 Shape                     // nil interface (type=nil, value=nil)
	var r1 *Rectangle                // nil pointer
	var s2 Shape = r1                // interface with nil concrete value

	fmt.Println("s1 == nil:", s1 == nil) // true
	fmt.Println("r1 == nil:", r1 == nil) // true
	fmt.Println("s2 == nil:", s2 == nil) // false! (type is *Rectangle)

	// 9. COMMON INTERFACES: fmt.Stringer

	fmt.Println("\n--- fmt.Stringer Interface ---")

	// fmt.Stringer: String() string
	// Used by fmt.Print functions

	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p) // calls p.String() automatically

	// 10. COMMON INTERFACES: error

	fmt.Println("\n--- error Interface ---")

	// error interface: Error() string

	err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	result, err := divideWithResult(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// 11. INTERFACE COMPOSITION

	fmt.Println("\n--- Interface Composition ---")

	// Combine interfaces
	type Reader interface {
		Read() string
	}

	type Writer interface {
		Write(string)
	}

	type ReadWriter interface {
		Reader
		Writer
	}

	// A type implementing ReadWriter must have both Read() and Write()

	// 12. CHECKING INTERFACE IMPLEMENTATION

	fmt.Println("\n--- Checking Implementation ---")

	var _ Shape = Rectangle{}  // compile-time check
	var _ Shape = (*Circle)(nil) // works for pointer receivers too

	// Check at runtime
	var unknown interface{} = Circle{5}

	if shape, ok := unknown.(Shape); ok {
		fmt.Printf("Implements Shape! Area: %.2f\n", shape.Area())
	}

	// 13. PRACTICAL: DEPENDENCY INJECTION

	fmt.Println("\n--- Dependency Injection ---")

	// Interface allows swapping implementations
	realDB := &RealDatabase{}
	mockDB := &MockDatabase{}

	service1 := UserService{db: realDB}
	service2 := UserService{db: mockDB}

	fmt.Println("Real:", service1.GetUser(1))
	fmt.Println("Mock:", service2.GetUser(1))
}

// METHODS FOR TYPES

// Rectangle methods
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle methods (Heron's formula)
func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// Square methods
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// HELPER FUNCTIONS

func printShapeInfo(s Shape) {
	fmt.Printf("  %T: Area=%.2f, Perimeter=%.2f\n", s, s.Area(), s.Perimeter())
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("  int: %d\n", v)
	case string:
		fmt.Printf("  string: %q\n", v)
	case float64:
		fmt.Printf("  float64: %.2f\n", v)
	case bool:
		fmt.Printf("  bool: %v\n", v)
	default:
		fmt.Printf("  unknown type: %T\n", v)
	}
}

// fmt.Stringer EXAMPLE

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// error INTERFACE EXAMPLE

type DivisionError struct {
	Dividend, Divisor float64
}

func (e DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by %.2f", e.Dividend, e.Divisor)
}

func divide(a, b float64) error {
	if b == 0 {
		return DivisionError{a, b}
	}
	return nil
}

func divideWithResult(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivisionError{a, b}
	}
	return a / b, nil
}

// DEPENDENCY INJECTION EXAMPLE

type Database interface {
	GetUser(id int) string
}

type RealDatabase struct{}

func (db *RealDatabase) GetUser(id int) string {
	return fmt.Sprintf("User %d from real database", id)
}

type MockDatabase struct{}

func (db *MockDatabase) GetUser(id int) string {
	return fmt.Sprintf("Mock user %d", id)
}

type UserService struct {
	db Database
}

func (s UserService) GetUser(id int) string {
	return s.db.GetUser(id)
}
