package main

import "fmt"

/*
  SLICES IN GO

KEY POINTS:
  - Slices are DYNAMIC, flexible views into arrays
  - A slice has 3 components: pointer, length, capacity
  - Slices are REFERENCE types: they point to an underlying array
  - Multiple slices can share the same underlying array
  - Zero value of a slice is nil

SYNTAX:
  var s []Type              // nil slice
  s := []Type{values}       // slice literal
  s := make([]Type, len)    // make with length
  s := make([]Type, len, cap) // make with length and capacity

SLICE VS ARRAY:
  - Array: [3]int  → fixed size, value type
  - Slice: []int   → dynamic size, reference type

*/

func main() {
	// 1. CREATING SLICES

	// Method 1: Slice literal
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice literal:", nums)

	// Method 2: From an array
	arr := [5]string{"a", "b", "c", "d", "e"}
	slice := arr[1:4] // elements at index 1, 2, 3
	fmt.Println("From array arr[1:4]:", slice)

	// Method 3: Using make
	s1 := make([]int, 3)    // length=3, capacity=3
	s2 := make([]int, 3, 5) // length=3, capacity=5
	fmt.Println("make([]int, 3):", s1, "len:", len(s1), "cap:", cap(s1))
	fmt.Println("make([]int, 3, 5):", s2, "len:", len(s2), "cap:", cap(s2))

	// Method 4: nil slice
	var nilSlice []int
	fmt.Println("nil slice:", nilSlice, "is nil:", nilSlice == nil)

	// 2. SLICE OPERATIONS: SLICING

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("\nOriginal:", data)
	fmt.Println("data[2:5]:", data[2:5])   // [2 3 4]
	fmt.Println("data[:3]:", data[:3])     // [0 1 2] (from start)
	fmt.Println("data[7:]:", data[7:])     // [7 8 9] (to end)
	fmt.Println("data[:]:", data[:])       // full slice (copy of reference)

	// 3. SLICES SHARE UNDERLYING ARRAY

	original := []int{1, 2, 3, 4, 5}
	sliceA := original[1:4] // [2 3 4]
	sliceB := original[2:5] // [3 4 5]

	fmt.Println("\n--- Shared underlying array ---")
	fmt.Println("original:", original)
	fmt.Println("sliceA:", sliceA)
	fmt.Println("sliceB:", sliceB)

	// Modify through sliceA
	sliceA[1] = 999

	fmt.Println("After sliceA[1] = 999:")
	fmt.Println("original:", original) // [1 2 999 4 5]
	fmt.Println("sliceA:", sliceA)     // [2 999 4]
	fmt.Println("sliceB:", sliceB)     // [999 4 5]

	// 4. LENGTH VS CAPACITY

	fmt.Println("\n--- Length vs Capacity ---")
	s := []int{1, 2, 3, 4, 5}
	printSlice("s", s)

	// Slice to zero length (capacity remains)
	s = s[:0]
	printSlice("s[:0]", s)

	// Extend length (up to capacity)
	s = s[:4]
	printSlice("s[:4]", s)

	// Drop first 2 elements (reduces capacity)
	s = s[2:]
	printSlice("s[2:]", s)

	// 5. APPENDING TO SLICES

	fmt.Println("\n--- Append ---")
	var growing []int
	printSlice("empty", growing)

	growing = append(growing, 1)
	printSlice("append 1", growing)

	growing = append(growing, 2, 3, 4)
	printSlice("append 2,3,4", growing)

	// Append another slice using ...
	more := []int{5, 6, 7}
	growing = append(growing, more...)
	printSlice("append slice", growing)

	// 6. CAPACITY GROWTH

	fmt.Println("\n--- Capacity growth ---")
	var capDemo []int
	prevCap := 0

	for i := 0; i < 20; i++ {
		capDemo = append(capDemo, i)
		if cap(capDemo) != prevCap {
			fmt.Printf("len=%2d, cap=%2d (grew from %d)\n", len(capDemo), cap(capDemo), prevCap)
			prevCap = cap(capDemo)
		}
	}

	// 7. COPYING SLICES

	fmt.Println("\n--- Copy ---")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))

	copied := copy(dst, src)
	fmt.Println("Source:", src)
	fmt.Println("Destination:", dst)
	fmt.Println("Elements copied:", copied)

	// Modifying dst doesn't affect src
	dst[0] = 999
	fmt.Println("After dst[0] = 999:")
	fmt.Println("Source:", src) // unchanged
	fmt.Println("Destination:", dst)

	// 8. ITERATING WITH RANGE

	fmt.Println("\n--- Range iteration ---")
	fruits := []string{"apple", "banana", "cherry"}

	// Index and value
	for i, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, fruit)
	}

	// Value only
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

	// 9. SLICES OF STRUCTS

	fmt.Println("\n--- Slice of structs ---")
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	for _, p := range people {
		fmt.Printf("%s is %d years old\n", p.Name, p.Age)
	}

	// 10. REMOVING ELEMENTS

	fmt.Println("\n--- Remove element ---")
	items := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Before:", items)

	// Remove element at index 2 ("c")
	i := 2
	items = append(items[:i], items[i+1:]...)
	fmt.Println("After removing index 2:", items)

	// 11. PREALLOCATING FOR PERFORMANCE

	fmt.Println("\n--- Preallocation ---")
	// Bad: grows capacity multiple times
	var slow []int
	for i := 0; i < 1000; i++ {
		slow = append(slow, i)
	}

	// Good: preallocate when size is known
	fast := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		fast = append(fast, i)
	}
	fmt.Println("Preallocated slice length:", len(fast))
}

func printSlice(name string, s []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n", name, len(s), cap(s), s)
}
