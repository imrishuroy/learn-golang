package main

import (
	"fmt"
	"math"
)

/*
  METHODS IN GO

KEY POINTS:
  - A method is a function with a special RECEIVER argument
  - The receiver appears between func keyword and method name
  - Methods can be defined on any type in the same package
  - Go doesn't have classes, but methods give types behavior

SYNTAX:
  func (r ReceiverType) MethodName(params) returnType { }

RECEIVER TYPES:
  1. Value receiver:   func (r Type) Method()
     - Receives a COPY of the value
     - Cannot modify the original
     - Safe for concurrent access

  2. Pointer receiver: func (r *Type) Method()
     - Receives a pointer to the value
     - CAN modify the original
     - More efficient for large structs

BEST PRACTICES:
  - If ANY method needs pointer receiver, ALL should use pointer
  - Use pointer receivers when: mutating state, large structs
  - Use value receivers when: small structs, read-only operations

*/

// TYPE DEFINITIONS

// Vertex represents a 2D point
type Vertex struct {
	X, Y float64
}

// Rectangle for area/perimeter examples
type Rectangle struct {
	Width, Height float64
}

// Circle for comparison with Rectangle
type Circle struct {
	Radius float64
}

// MyInt is a custom type based on int
type MyInt int

// Counter demonstrates state modification
type Counter struct {
	value int
}

func main() {
	// 1. BASIC METHOD - VALUE RECEIVER

	fmt.Println("--- Value Receiver ---")

	v := Vertex{3, 4}
	fmt.Println("Vertex:", v)
	fmt.Println("Magnitude:", v.Magnitude()) // method call

	// Compare with regular function
	fmt.Println("MagnitudeFunc:", MagnitudeFunc(v))

	// 2. POINTER RECEIVER - MODIFYING STATE

	fmt.Println("\n--- Pointer Receiver ---")

	v2 := Vertex{3, 4}
	fmt.Println("Before Scale:", v2)

	v2.Scale(2) // modifies v2
	fmt.Println("After Scale(2):", v2)

	// Go auto-dereferences: these are equivalent
	v3 := Vertex{1, 2}
	ptr := &v3
	ptr.Scale(10) // (*ptr).Scale(10)
	fmt.Println("Scaled via pointer:", v3)

	// 3. VALUE VS POINTER RECEIVERS

	fmt.Println("\n--- Value vs Pointer ---")

	v4 := Vertex{5, 5}

	// Value receiver - operates on copy
	v4.TryScaleValue(100)
	fmt.Println("After TryScaleValue:", v4) // unchanged!

	// Pointer receiver - operates on original
	v4.Scale(100)
	fmt.Println("After Scale:", v4) // changed!

	// 4. METHODS ON CUSTOM TYPES

	fmt.Println("\n--- Methods on Custom Types ---")

	// Methods can be defined on ANY type (not just structs)
	num := MyInt(-5)
	fmt.Println("MyInt:", num)
	fmt.Println("Abs:", num.Abs())
	fmt.Println("Double:", num.Double())

	// 5. PRACTICAL: SHAPE METHODS

	fmt.Println("\n--- Shape Methods ---")

	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	fmt.Println("Rectangle:", rect)
	fmt.Printf("  Area: %.2f\n", rect.Area())
	fmt.Printf("  Perimeter: %.2f\n", rect.Perimeter())

	fmt.Println("Circle:", circle)
	fmt.Printf("  Area: %.2f\n", circle.Area())
	fmt.Printf("  Circumference: %.2f\n", circle.Circumference())

	// 6. METHOD CHAINING

	fmt.Println("\n--- Method Chaining ---")

	v5 := Vertex{1, 1}
	fmt.Println("Original:", v5)

	// Chain methods that return *Vertex
	v5.ScaleChain(2).ScaleChain(3).TranslateChain(10, 10)
	fmt.Println("After chaining:", v5)

	// 7. METHODS WITH STATE

	fmt.Println("\n--- Methods with State ---")

	c := Counter{value: 0}

	c.Increment()
	c.Increment()
	c.Increment()
	fmt.Println("After 3 increments:", c.Value())

	c.Add(10)
	fmt.Println("After Add(10):", c.Value())

	c.Reset()
	fmt.Println("After Reset:", c.Value())

	// 8. NIL RECEIVERS

	fmt.Println("\n--- Nil Receivers ---")

	// Methods can handle nil receivers gracefully
	var nilVertex *Vertex
	fmt.Println("Nil vertex safe method:", nilVertex.SafeMagnitude())

	// This would panic: fmt.Println(nilVertex.Magnitude())

	// 9. EMBEDDING AND METHOD PROMOTION

	fmt.Println("\n--- Embedding ---")

	type ColoredVertex struct {
		Vertex // embedded - methods are promoted
		Color  string
	}

	cv := ColoredVertex{
		Vertex: Vertex{3, 4},
		Color:  "red",
	}

	// Can call Vertex methods directly
	fmt.Println("ColoredVertex:", cv)
	fmt.Println("Magnitude:", cv.Magnitude()) // promoted method
	fmt.Println("Color:", cv.Color)

	// 10. METHODS AS VALUES

	fmt.Println("\n--- Methods as Values ---")

	v6 := Vertex{3, 4}

	// Method value - bound to specific receiver
	magFunc := v6.Magnitude
	fmt.Println("Method value:", magFunc())

	// Method expression - takes receiver as first arg
	magExpr := Vertex.Magnitude
	fmt.Println("Method expression:", magExpr(v6))

	// Useful for passing to higher-order functions
	vertices := []Vertex{{3, 4}, {5, 12}, {8, 15}}
	for _, v := range vertices {
		fmt.Printf("  Vertex%v magnitude: %.2f\n", v, v.Magnitude())
	}
}

// METHODS

// Magnitude returns the length of the vector (value receiver)
func (v Vertex) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MagnitudeFunc is the same as Magnitude but as a regular function
func MagnitudeFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale multiplies both coordinates (pointer receiver - modifies)
func (v *Vertex) Scale(factor float64) {
	v.X *= factor
	v.Y *= factor
}

// TryScaleValue tries to scale but can't (value receiver)
func (v Vertex) TryScaleValue(factor float64) {
	v.X *= factor // only modifies the copy!
	v.Y *= factor
}

// ScaleChain scales and returns pointer for chaining
func (v *Vertex) ScaleChain(factor float64) *Vertex {
	v.X *= factor
	v.Y *= factor
	return v
}

// TranslateChain moves the vertex and returns pointer for chaining
func (v *Vertex) TranslateChain(dx, dy float64) *Vertex {
	v.X += dx
	v.Y += dy
	return v
}

// SafeMagnitude handles nil receiver
func (v *Vertex) SafeMagnitude() float64 {
	if v == nil {
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods on custom type MyInt
func (n MyInt) Abs() MyInt {
	if n < 0 {
		return -n
	}
	return n
}

func (n MyInt) Double() MyInt {
	return n * 2
}

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

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// Counter methods
func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Add(n int) {
	c.value += n
}

func (c *Counter) Reset() {
	c.value = 0
}

func (c Counter) Value() int {
	return c.value
}
