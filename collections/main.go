package main

import (
	"fmt"
	"sort"
)

/*
  COLLECTIONS IN GO: ARRAYS, SLICES, AND MAPS

KEY POINTS:
  Arrays: Fixed size, value type, size is part of type
  Slices: Dynamic size, reference type, built on arrays
  Maps: Key-value pairs, reference type, unordered

ARRAYS:
  var arr [5]int          // Zero-valued array
  arr := [5]int{1,2,3,4,5} // Array literal
  arr := [...]int{1,2,3}   // Size inferred

SLICES:
  var s []int             // Nil slice
  s := []int{1,2,3}       // Slice literal
  s := make([]int, len)   // Make with length
  s := make([]int, len, cap) // Make with capacity

MAPS:
  var m map[string]int    // Nil map (can't write!)
  m := make(map[string]int) // Empty map
  m := map[string]int{"a": 1} // Map literal

*/

func main() {
	// ==================== ARRAYS ====================

	fmt.Println("========== ARRAYS ==========")

	// 1. Array Declaration
	fmt.Println("\n--- Array Declaration ---")

	var arr1 [5]int                    // Zero-valued
	arr2 := [5]int{1, 2, 3, 4, 5}      // Literal
	arr3 := [...]int{10, 20, 30}       // Size inferred
	arr4 := [5]int{0: 100, 4: 500}     // Specific indices

	fmt.Println("Zero-valued:", arr1)
	fmt.Println("Literal:", arr2)
	fmt.Println("Inferred size:", arr3, "length:", len(arr3))
	fmt.Println("Specific indices:", arr4)

	// 2. Arrays are Value Types (copying)
	fmt.Println("\n--- Arrays are Value Types ---")

	original := [3]string{"a", "b", "c"}
	copied := original // Full copy!
	copied[0] = "X"

	fmt.Println("Original:", original) // [a b c] - unchanged
	fmt.Println("Copied:", copied)     // [X b c]

	// 3. Iterating Arrays
	fmt.Println("\n--- Iterating Arrays ---")

	nums := [5]int{10, 20, 30, 40, 50}

	fmt.Print("Index only: ")
	for i := range nums {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Print("Index and value: ")
	for i, v := range nums {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	// ==================== SLICES ====================

	fmt.Println("\n========== SLICES ==========")

	// 4. Slice Creation
	fmt.Println("\n--- Slice Creation ---")

	var nilSlice []int                    // Nil slice
	emptySlice := []int{}                 // Empty slice (not nil)
	literalSlice := []int{1, 2, 3, 4, 5}  // Slice literal
	madeSlice := make([]int, 3)           // Length 3, cap 3
	madeSliceCap := make([]int, 3, 10)    // Length 3, cap 10

	fmt.Printf("Nil slice: %v, nil=%v\n", nilSlice, nilSlice == nil)
	fmt.Printf("Empty slice: %v, nil=%v\n", emptySlice, emptySlice == nil)
	fmt.Println("Literal slice:", literalSlice)
	fmt.Printf("make([]int, 3): %v, len=%d, cap=%d\n", madeSlice, len(madeSlice), cap(madeSlice))
	fmt.Printf("make([]int, 3, 10): %v, len=%d, cap=%d\n", madeSliceCap, len(madeSliceCap), cap(madeSliceCap))

	// 5. Slicing Operations
	fmt.Println("\n--- Slicing Operations ---")

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original:", data)
	fmt.Println("data[2:5]:", data[2:5])   // [2 3 4]
	fmt.Println("data[:3]:", data[:3])     // [0 1 2]
	fmt.Println("data[7:]:", data[7:])     // [7 8 9]
	fmt.Println("data[:]:", data[:])       // Full copy reference

	// 6. Slices Share Underlying Array
	fmt.Println("\n--- Slices Share Array ---")

	base := []int{1, 2, 3, 4, 5}
	slice1 := base[1:4]
	slice2 := base[2:5]

	fmt.Println("base:", base)
	fmt.Println("slice1 (base[1:4]):", slice1)
	fmt.Println("slice2 (base[2:5]):", slice2)

	slice1[1] = 999 // Modifies shared array

	fmt.Println("After slice1[1] = 999:")
	fmt.Println("base:", base)     // [1 2 999 4 5]
	fmt.Println("slice1:", slice1) // [2 999 4]
	fmt.Println("slice2:", slice2) // [999 4 5]

	// 7. Append
	fmt.Println("\n--- Append ---")

	var growing []int
	fmt.Printf("Initial: len=%d, cap=%d\n", len(growing), cap(growing))

	growing = append(growing, 1)
	fmt.Printf("After append(1): len=%d, cap=%d, %v\n", len(growing), cap(growing), growing)

	growing = append(growing, 2, 3, 4)
	fmt.Printf("After append(2,3,4): len=%d, cap=%d, %v\n", len(growing), cap(growing), growing)

	more := []int{5, 6, 7}
	growing = append(growing, more...) // Append slice
	fmt.Printf("After append(slice...): len=%d, cap=%d, %v\n", len(growing), cap(growing), growing)

	// 8. Copy
	fmt.Println("\n--- Copy ---")

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src)

	fmt.Println("Source:", src)
	fmt.Println("Destination:", dst)
	fmt.Println("Elements copied:", n)

	dst[0] = 999
	fmt.Println("After dst[0] = 999:")
	fmt.Println("Source:", src) // Unchanged
	fmt.Println("Destination:", dst)

	// 9. Remove Element
	fmt.Println("\n--- Remove Element ---")

	items := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Before:", items)

	i := 2 // Remove "c"
	items = append(items[:i], items[i+1:]...)
	fmt.Println("After removing index 2:", items)

	// ==================== MAPS ====================

	fmt.Println("\n========== MAPS ==========")

	// 10. Map Creation
	fmt.Println("\n--- Map Creation ---")

	var nilMap map[string]int              // Nil map
	emptyMap := map[string]int{}           // Empty map
	madeMap := make(map[string]int)        // Empty map with make
	literalMap := map[string]int{          // Map literal
		"alice": 95,
		"bob":   87,
	}

	fmt.Printf("Nil map: %v, nil=%v\n", nilMap, nilMap == nil)
	fmt.Printf("Empty map: %v\n", emptyMap)
	fmt.Println("Made map:", madeMap)
	fmt.Println("Literal map:", literalMap)

	// 11. Map Operations
	fmt.Println("\n--- Map Operations ---")

	scores := make(map[string]int)

	// Insert/Update
	scores["alice"] = 95
	scores["bob"] = 87
	scores["charlie"] = 92
	fmt.Println("After inserts:", scores)

	// Read
	fmt.Println("alice's score:", scores["alice"])
	fmt.Println("unknown's score:", scores["unknown"]) // 0 (zero value)

	// Check existence (comma-ok idiom)
	val, exists := scores["alice"]
	fmt.Printf("alice: value=%d, exists=%v\n", val, exists)

	val, exists = scores["unknown"]
	fmt.Printf("unknown: value=%d, exists=%v\n", val, exists)

	// Delete
	delete(scores, "bob")
	fmt.Println("After delete(bob):", scores)

	// 12. Iterating Maps
	fmt.Println("\n--- Iterating Maps ---")

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	fmt.Println("Key-value pairs:")
	for key, value := range colors {
		fmt.Printf("  %s: %s\n", key, value)
	}

	fmt.Print("Keys only: ")
	for key := range colors {
		fmt.Print(key, " ")
	}
	fmt.Println()

	// 13. Map as Set
	fmt.Println("\n--- Map as Set ---")

	visited := make(map[string]bool)
	visited["Paris"] = true
	visited["London"] = true
	visited["Tokyo"] = true

	cities := []string{"Paris", "NYC", "Tokyo"}
	for _, city := range cities {
		if visited[city] {
			fmt.Printf("%s: visited\n", city)
		} else {
			fmt.Printf("%s: not visited\n", city)
		}
	}

	// 14. Counting with Maps
	fmt.Println("\n--- Counting ---")

	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	count := make(map[string]int)

	for _, word := range words {
		count[word]++ // Zero value (0) + 1 works!
	}

	fmt.Println("Word counts:")
	for word, n := range count {
		fmt.Printf("  %s: %d\n", word, n)
	}

	// 15. Sorting Map Keys
	fmt.Println("\n--- Sorting Map Keys ---")

	ages := map[string]int{"charlie": 35, "alice": 30, "bob": 25}

	// Get sorted keys
	keys := make([]string, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Sorted by name:")
	for _, k := range keys {
		fmt.Printf("  %s: %d\n", k, ages[k])
	}
}
