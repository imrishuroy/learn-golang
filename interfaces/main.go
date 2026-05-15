package main

import (
	"fmt"
	"math"
)

/*
  INTERFACES IN GO

KEY POINTS:
  - An interface defines a set of method signatures
  - Types IMPLICITLY implement interfaces (no "implements" keyword)
  - If a type has all the methods, it implements the interface
  - Interfaces enable polymorphism and decoupling

SYNTAX:
  type InterfaceName interface {
      MethodName(params) returnType
  }

WHY INTERFACES:
  1. Polymorphism: Different types, same behavior
  2. Decoupling: Code depends on behavior, not concrete types
  3. Testing: Easy to create mocks
  4. Flexibility: Swap implementations without changing code

BEST PRACTICES:
  - Keep interfaces small (1-3 methods)
  - Define interfaces where they're USED, not implemented
  - "Accept interfaces, return concrete types"

COMMON INTERFACES:
  - fmt.Stringer: String() string
  - error: Error() string
  - io.Reader: Read([]byte) (int, error)
  - io.Writer: Write([]byte) (int, error)

*/

// 1. DEFINE AN INTERFACE

// Session defines behavior for charging sessions
type Session interface {
	Start()
	Stop()
	GetStatus() string
}

// Shape defines behavior for geometric shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Describer adds description capability
type Describer interface {
	Describe() string
}

// 2. TYPES THAT IMPLEMENT INTERFACES

// EvChargeSession implements Session
type EvChargeSession struct {
	Id      string
	Watts   int
	Active  bool
}

func (e *EvChargeSession) Start() {
	e.Active = true
	fmt.Printf("EV Session %s started at %dW\n", e.Id, e.Watts)
}

func (e *EvChargeSession) Stop() {
	e.Active = false
	fmt.Printf("EV Session %s stopped\n", e.Id)
}

func (e *EvChargeSession) GetStatus() string {
	if e.Active {
		return "charging"
	}
	return "idle"
}

// HybridChargeSession also implements Session
type HybridChargeSession struct {
	Id       string
	Watts    int
	GasLevel int
	Active   bool
}

func (h *HybridChargeSession) Start() {
	h.Active = true
	fmt.Printf("Hybrid Session %s started at %dW (Gas: %d%%)\n", h.Id, h.Watts, h.GasLevel)
}

func (h *HybridChargeSession) Stop() {
	h.Active = false
	fmt.Printf("Hybrid Session %s stopped\n", h.Id)
}

func (h *HybridChargeSession) GetStatus() string {
	if h.Active {
		return "hybrid-charging"
	}
	return "idle"
}

// Rectangle implements Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 3. USING INTERFACES

	fmt.Println("--- Using Interfaces ---")

	// Both types can be used as Session
	ev := &EvChargeSession{Id: "EV-001", Watts: 50}
	hybrid := &HybridChargeSession{Id: "HYB-001", Watts: 22, GasLevel: 75}

	// Function accepts any Session
	startSession(ev)
	startSession(hybrid)

	fmt.Println("EV Status:", ev.GetStatus())
	fmt.Println("Hybrid Status:", hybrid.GetStatus())

	// 4. INTERFACE VALUES

	fmt.Println("\n--- Interface Values ---")

	var session Session // Interface variable

	session = ev
	fmt.Printf("Interface holds: %T\n", session)
	session.Stop()

	session = hybrid
	fmt.Printf("Interface holds: %T\n", session)
	session.Stop()

	// 5. POLYMORPHISM WITH SHAPES

	fmt.Println("\n--- Polymorphism ---")

	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
		Rectangle{Width: 3, Height: 3},
	}

	for _, shape := range shapes {
		fmt.Printf("%T: Area=%.2f, Perimeter=%.2f\n",
			shape, shape.Area(), shape.Perimeter())
	}

	// Calculate total area
	total := totalArea(shapes)
	fmt.Printf("Total area: %.2f\n", total)

	// 6. EMPTY INTERFACE (any)

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
	printAll(1, "two", 3.0, true, []int{1, 2, 3})

	// 7. TYPE ASSERTIONS

	fmt.Println("\n--- Type Assertions ---")

	var i interface{} = "hello"

	// Safe type assertion with ok check
	str, ok := i.(string)
	fmt.Printf("String assertion: value=%q, ok=%v\n", str, ok)

	num, ok := i.(int)
	fmt.Printf("Int assertion: value=%d, ok=%v\n", num, ok)

	// 8. TYPE SWITCH

	fmt.Println("\n--- Type Switch ---")

	describe(42)
	describe("hello")
	describe(3.14)
	describe(Rectangle{5, 3})
	describe(true)

	// 9. STRINGER INTERFACE

	fmt.Println("\n--- Stringer Interface ---")

	// fmt.Stringer: String() string
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p) // Calls p.String() automatically

	// 10. ERROR INTERFACE

	fmt.Println("\n--- Error Interface ---")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 11. INTERFACE COMPOSITION

	fmt.Println("\n--- Interface Composition ---")

	// Combine interfaces
	type ReadWriter interface {
		Reader
		Writer
	}

	// A type implementing ReadWriter must have both Read() and Write()
	fmt.Println("Interfaces can be composed from other interfaces")

	// 12. CHECKING INTERFACE IMPLEMENTATION

	fmt.Println("\n--- Checking Implementation ---")

	// Compile-time check that types implement interfaces
	var _ Session = (*EvChargeSession)(nil)
	var _ Session = (*HybridChargeSession)(nil)
	var _ Shape = Rectangle{}
	var _ Shape = Circle{}

	fmt.Println("All types correctly implement their interfaces")
}

// HELPER FUNCTIONS

// startSession accepts any Session implementation
func startSession(s Session) {
	s.Start()
}

// totalArea calculates total area of all shapes
func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// printAll prints any number of any type
func printAll(args ...interface{}) {
	for i, arg := range args {
		fmt.Printf("  [%d] %v (type: %T)\n", i, arg, arg)
	}
}

// describe uses type switch to handle different types
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("  int: %d\n", v)
	case string:
		fmt.Printf("  string: %q\n", v)
	case float64:
		fmt.Printf("  float64: %.2f\n", v)
	case Shape:
		fmt.Printf("  Shape: area=%.2f\n", v.Area())
	default:
		fmt.Printf("  unknown: %T\n", v)
	}
}

// STRINGER EXAMPLE

type Person struct {
	Name string
	Age  int
}

// String implements fmt.Stringer
func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// ERROR EXAMPLE

type DivisionError struct {
	Dividend, Divisor float64
}

func (e DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %.0f by %.0f", e.Dividend, e.Divisor)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivisionError{a, b}
	}
	return a / b, nil
}

// INTERFACE COMPOSITION TYPES

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}
