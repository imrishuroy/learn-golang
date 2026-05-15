package main

import "fmt"

// Counter demonstrates pointer vs value receivers
type Counter struct {
	count int
}

/*
  POINTERS IN GO

KEY POINTS:
  - A pointer stores the memory address of another variable
  - Pointers let you share data without copying
  - Zero value of a pointer is nil

OPERATORS:
  &  (address-of)   → gets the memory address of a variable
  *  (dereference)  → accesses the value at a memory address

SYNTAX:
  var p *int       // p is a pointer to int (currently nil)
  p = &x           // p now holds address of x
  *p = 10          // set value at address to 10
  y := *p          // read value at address

WHY USE POINTERS:
  1. Avoid copying large structs (performance)
  2. Allow functions to modify caller's variables
  3. Share data between different parts of program
  4. Implement data structures (linked lists, trees)

*/

func main() {
	// 1. BASIC POINTER USAGE

	x := 42
	p := &x // p points to x

	fmt.Println("--- Basic Usage ---")
	fmt.Println("x =", x)           // 42
	fmt.Println("&x =", &x)         // 0xc000... (memory address)
	fmt.Println("p =", p)           // same address
	fmt.Println("*p =", *p)         // 42 (value at address)

	// Modify x through the pointer
	*p = 100
	fmt.Println("After *p = 100:")
	fmt.Println("x =", x)           // 100 (changed!)
	fmt.Println("*p =", *p)         // 100

	// 2. POINTER TO POINTER

	fmt.Println("\n--- Pointer to Pointer ---")
	a := 10
	pa := &a   // pointer to a
	ppa := &pa // pointer to pointer

	fmt.Println("a =", a)
	fmt.Println("*pa =", *pa)
	fmt.Println("**ppa =", **ppa)

	**ppa = 20
	fmt.Println("After **ppa = 20, a =", a)

	// 3. NIL POINTERS

	fmt.Println("\n--- Nil Pointers ---")
	var nilPtr *int
	fmt.Println("nil pointer:", nilPtr)
	fmt.Println("is nil:", nilPtr == nil)

	// Demonstrate safe nil check with a function
	safePrint(nilPtr)      // handles nil
	safePrint(&x)          // prints value

	// 4. FUNCTIONS WITH POINTERS

	fmt.Println("\n--- Functions with Pointers ---")

	// Pass by value (copy)
	num := 5
	fmt.Println("Before doubleValue:", num)
	doubleValue(num)
	fmt.Println("After doubleValue:", num) // unchanged

	// Pass by pointer (reference)
	fmt.Println("Before doublePointer:", num)
	doublePointer(&num)
	fmt.Println("After doublePointer:", num) // doubled!

	// 5. RETURNING POINTERS FROM FUNCTIONS

	fmt.Println("\n--- Returning Pointers ---")

	// This is safe in Go - the variable escapes to heap
	ptr := createInt(42)
	fmt.Println("Created int:", *ptr)

	// Contrast with C/C++ where this would be undefined behavior

	// 6. POINTERS TO STRUCTS

	fmt.Println("\n--- Pointers to Structs ---")

	type Person struct {
		Name string
		Age  int
	}

	alice := Person{"Alice", 30}
	pAlice := &alice

	// Both work - Go auto-dereferences for struct fields
	fmt.Println("(*pAlice).Name:", (*pAlice).Name)
	fmt.Println("pAlice.Name:", pAlice.Name) // shorthand

	// Modify through pointer
	pAlice.Age = 31
	fmt.Println("After birthday:", alice)

	// 7. NEW() FUNCTION

	fmt.Println("\n--- new() Function ---")

	// new() allocates zeroed memory and returns a pointer
	intPtr := new(int) // *int pointing to 0
	fmt.Println("new(int):", *intPtr)

	*intPtr = 42
	fmt.Println("After assignment:", *intPtr)

	// Equivalent to:
	// var i int
	// intPtr := &i

	type Point struct {
		X, Y int
	}

	pointPtr := new(Point) // *Point with X=0, Y=0
	fmt.Println("new(Point):", *pointPtr)

	// 8. POINTER RECEIVERS VS VALUE RECEIVERS

	fmt.Println("\n--- Pointer vs Value Receivers ---")

	c := Counter{count: 0}

	// Value receiver - doesn't modify original
	c.incrementValue()
	fmt.Println("After incrementValue:", c.count) // still 0

	// Pointer receiver - modifies original
	c.incrementPointer()
	fmt.Println("After incrementPointer:", c.count) // 1

	// 9. SLICES AND MAPS ARE ALREADY REFERENCES

	fmt.Println("\n--- Slices/Maps (already references) ---")

	// No need to pass pointer - slice header contains pointer
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println("Slice after modification:", slice) // [100 2 3]

	// But append can't change caller's slice without pointer
	slice2 := []int{1, 2, 3}
	appendToSlice(slice2)
	fmt.Println("Slice after append (no ptr):", slice2) // [1 2 3] unchanged

	appendToSlicePtr(&slice2)
	fmt.Println("Slice after append (ptr):", slice2) // [1 2 3 4]

	// 10. PRACTICAL EXAMPLE: SWAPPING VALUES

	fmt.Println("\n--- Swap Example ---")

	i, j := 10, 20
	fmt.Printf("Before swap: i=%d, j=%d\n", i, j)

	swap(&i, &j)
	fmt.Printf("After swap: i=%d, j=%d\n", i, j)
}

// Pass by value - receives a copy
func doubleValue(n int) {
	n = n * 2 // only modifies the copy
}

// Pass by pointer - receives address of original
func doublePointer(n *int) {
	*n = *n * 2 // modifies the original
}

// Returns pointer to local variable (safe in Go)
func createInt(val int) *int {
	result := val
	return &result // result escapes to heap
}

// Value receiver - operates on a copy
func (c Counter) incrementValue() {
	c.count++ // modifies copy, not original
}

// Pointer receiver - operates on original
func (c *Counter) incrementPointer() {
	c.count++ // modifies original
}

// Slices pass underlying array by reference
func modifySlice(s []int) {
	s[0] = 100 // modifies original array
}

// But can't change slice header (len, cap, ptr) without pointer
func appendToSlice(s []int) {
	s = append(s, 4) // local change only
}

func appendToSlicePtr(s *[]int) {
	*s = append(*s, 4) // modifies caller's slice
}

// Classic swap using pointers
func swap(a, b *int) {
	*a, *b = *b, *a
}

// safePrint demonstrates nil checking before dereferencing
func safePrint(p *int) {
	if p != nil {
		fmt.Println("Value:", *p)
	} else {
		fmt.Println("Cannot dereference nil pointer")
	}
}
