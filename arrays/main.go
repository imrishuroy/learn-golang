package main

import "fmt"

/*
  ARRAYS IN GO

KEY POINTS:
  - Arrays have a FIXED size defined at compile time
  - Size is part of the type: [3]int and [5]int are different types
  - Arrays are VALUE types: assignment/passing creates a copy
  - Zero value: all elements initialized to their zero value

SYNTAX:
  var arr [size]Type           // declaration
  arr := [size]Type{values}    // literal
  arr := [...]Type{values}     // compiler counts size

WHEN TO USE:
  - When you know the exact size at compile time
  - When you need value semantics (copying)
  - In practice, slices are used more often for flexibility

*/

func main() {
	// 1. DECLARING ARRAYS

	// Method 1: Declare then assign
	var a [3]int // zero value: [0 0 0]
	a[0] = 10
	a[1] = 20
	a[2] = 30
	fmt.Println("Method 1 - declare then assign:", a)

	// Method 2: Array literal
	b := [3]int{100, 200, 300}
	fmt.Println("Method 2 - array literal:", b)

	// Method 3: Let compiler count with [...]
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Method 3 - compiler counts:", c, "length:", len(c))

	// Method 4: Initialize specific indices
	d := [5]int{0: 10, 4: 50} // index 0 = 10, index 4 = 50, rest = 0
	fmt.Println("Method 4 - specific indices:", d)

	// 2. ACCESSING ELEMENTS

	names := [3]string{"Alice", "Bob", "Charlie"}

	fmt.Println("\nFirst element:", names[0])
	fmt.Println("Last element:", names[len(names)-1])

	// Modify an element
	names[1] = "Bobby"
	fmt.Println("After modification:", names)

	// 3. ITERATING OVER ARRAYS

	nums := [4]int{10, 20, 30, 40}

	// Method 1: Traditional for loop
	fmt.Print("\nTraditional loop: ")
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()

	// Method 2: Range (preferred)
	fmt.Print("Range loop: ")
	for index, value := range nums {
		fmt.Printf("[%d]=%d ", index, value)
	}
	fmt.Println()

	// Method 3: Range - value only (ignore index with _)
	fmt.Print("Values only: ")
	for _, value := range nums {
		fmt.Print(value, " ")
	}
	fmt.Println()

	// 4. ARRAYS ARE VALUE TYPES (COPYING)

	original := [3]int{1, 2, 3}
	copied := original // creates a FULL COPY

	copied[0] = 999

	fmt.Println("\nOriginal:", original) // [1 2 3] - unchanged
	fmt.Println("Copied:", copied)       // [999 2 3] - only copy changed

	// 5. MULTIDIMENSIONAL ARRAYS

	// 2D array (3 rows x 2 columns)
	matrix := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println("\n2D Array:")
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d] = %d  ", i, j, val)
		}
		fmt.Println()
	}

	// 6. COMPARING ARRAYS

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	fmt.Println("\narr1 == arr2:", arr1 == arr2) // true
	fmt.Println("arr1 == arr3:", arr1 == arr3)   // false

	// 7. PASSING ARRAYS TO FUNCTIONS

	arr := [3]int{10, 20, 30}
	fmt.Println("\nBefore function call:", arr)

	modifyArray(arr) // passes a COPY
	fmt.Println("After modifyArray (by value):", arr) // unchanged

	modifyArrayPtr(&arr) // passes pointer
	fmt.Println("After modifyArrayPtr (by pointer):", arr) // changed

	// 8. ARRAY LENGTH AND CAPACITY

	fixed := [5]int{1, 2, 3, 4, 5}
	fmt.Println("\nArray length:", len(fixed)) // 5
	// Note: Arrays don't have cap() - their capacity equals length
}

// modifyArray receives a COPY of the array
func modifyArray(arr [3]int) {
	arr[0] = 999 // modifies the copy, not the original
}

// modifyArrayPtr receives a pointer to the array
func modifyArrayPtr(arr *[3]int) {
	arr[0] = 999 // modifies the original
}
