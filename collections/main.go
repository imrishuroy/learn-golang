package main

import "fmt"

func main() {
	// collection in Go - Arrays, slices & maps

	// Array

	/*
		Arrays are a sequence of items of a single type. An array cannot be resized.
		Also, you have to explicitly define the length of an array in Go.
		That’s part of the type of an array, meaning together the size and the element
		type form the type of the array. Another caveat is that you cannot use a variable
		to set the length of the array.
	*/

	var n [2]string
	n[0] = "Rishu"
	n[1] = "Kumar"

	fmt.Println("Array", n)

	n2 := [2]string{"Prince", "Roy"}
	fmt.Println("Array 2", n2)

	nums := [6]int{2, 4, 6, 8, 10} // default value is 0
	fmt.Println("Nums array", nums)

	// In this we omitted the size part of the type with three ellipses ( 3 dots )
	// and the size is inferred
	nums2 := [...]int{21, 31, 41, 51, 61, 70}
	fmt.Println(nums2)
	// nums2[8] = 2 -> will give runtime error
	fmt.Println("length of of nums2", len(nums2))

	// ** Arrays are value types. This means passing an array to a function,
	//	or returning it from a function, creates a copy of the original array.

	n3 := [2]string{"Rishu", "Kumar"}

	n4 := n3
	n4[1] = "Roy"

	fmt.Println("n3", n3)
	fmt.Println("n4", n4)

	// looping through the array
	nums3 := [6]int{6, 9, 3, 41, 33, 55}
	fmt.Println(nums3)

	for i := range nums3 { // only getting the index
		fmt.Print(i, " ")
	}
	fmt.Println("")

	for i, item := range nums3 { // getting index and item
		fmt.Print(i, ": ", item, " ")
	}
	fmt.Println("")

	for i := 0; i < len(nums3); i++ { // older version
		fmt.Print(nums3[i], " ")
	}
	fmt.Println("")

	// Slice
	/*
		This header has pointer to the actual values array, and two additional values
		containing actual length and current capacity. With this structure it should be evident
		how a slice can dynamically grow as needed.
	*/

	/*
		ptr is a pointer to the actual data.
		len is the current length of the items identified by the ptr element.
		cap is the capacity of the slice, how many elements can be stored
		before it needs to be made larger.
	*/

	var sliceVar []int // slices declared with var are uninitialized, and are know as nill slices
	fmt.Println(len(sliceVar), cap(sliceVar))

	// add item to a slice
	intSlice := []int{1, 2, 3}
	fmt.Println("before append", len(intSlice))
	intSlice = append(intSlice, 10) // an append created a new slice
	fmt.Println("after appenf", len(intSlice))

	s1 := []int{1, 2, 3}
	s2 := []int{6, 7, 8, 9}

	s3 := copy(s1, s2)
	fmt.Println(s3, s1, s2)

	// create slice from array
	carMakers := [4]string{"Tesla", "Ford", "Volkswagen", "Mercedes"}
	carMakersSlice := carMakers[2:4] // 2 is inclusive and 4 is exclusive
	// [2:]

	fmt.Println(carMakersSlice)

	// Map
	// key value pairs
	// the key as the identifier that is used to find an associate value
	// With arrays or slices we retrieve using an index, with maps you
	// retrieve using the key

	chargeSessions := map[string]int{
		"111": 420, "222": 395, "333": 456,
	}
	key := "111"
	fmt.Println("value for", key, "is", chargeSessions[key])

	// map initialization

	cs := make(map[string]int) // Declare and initialize a map with 'make'
	cs["11111"] = 420
	fmt.Println(cs)

	// Declare an empty map with a literal.
	nilMap := map[int]string{} // this is identical to using make with {}
	fmt.Println(nilMap)

	// Declare and populate a map with a literal.
	// Keys separated from values with a colon,
	// and commas separate key-value pairs.
	m := map[string]int{

		"33333": 343, "44444": 410,
	}
	fmt.Println(m)

	var m2 map[string]int
	if m2 == nil {
		fmt.Println("Map is nil, initializing")
		m2 = map[string]int{}
	}

	m2["22222"] = 389 // Assignment succeeds after the map has been initialized.
	fmt.Println(m2)   // without the initialization the assign would cause a panic

	/*
		Useful is when a map that has values that are of the type bool and is queried
		for a key that does not exist, it returns a bool of false. This can be useful
		when building logic in a program. Here is how this could work, the map is
		queried for three values, one that does not exist.
	*/

	cs1 := map[string]bool{"111": true, "333": true}

	for _, id := range []string{"111", "222", "333"} {
		if cs1[id] {
			fmt.Printf("%s has a charge session.\n", id)
		} else {
			fmt.Printf("%s does not have a charge session.\n", id)
		}
	}

	// For an integer this works similar, only if it doesn't exists it return a zero (0)
	cs4 := map[string]int{"11111": 399,
		"33333": 400,
	}
	id := "22222"
	if cs4[id] != 0 {
		fmt.Printf("%s has a charge session.\n", id)
	} else {
		fmt.Printf("%s does not have a charge session.\n", id)
	}

	// This can also be done with user defined types. This example will show,
	// that the fields in the type definition get returned with their defaults.

	// delete
	map1 := map[string]int{"item1": 0, "item2": 1}
	delete(map1, "item2") // if we don't write the correct key, it don't do anything
	fmt.Println(map1)

	// loop

	map2 := map[string]int{"item1": 0, "item2": 1, "item3": 2}
	for key, value := range map2 {
		fmt.Println(key, value)

	}

	lenMap := len(map2)
	fmt.Println(lenMap)

}
