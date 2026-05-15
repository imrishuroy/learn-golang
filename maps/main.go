package main

import "fmt"

/*
  MAPS IN GO

KEY POINTS:
  - Maps are key-value pairs (like dict in Python, HashMap in Java)
  - Maps are REFERENCE types: zero value is nil
  - Keys must be comparable (==, !=): no slices, maps, or functions
  - Values can be any type
  - Map access is NOT thread-safe (use sync.Map or mutex)
  - Iteration order is NOT guaranteed

SYNTAX:
  var m map[KeyType]ValueType       // nil map (can't write to it!)
  m := make(map[KeyType]ValueType)  // empty, usable map
  m := map[KeyType]ValueType{}      // empty map literal
  m := map[KeyType]ValueType{k: v}  // map literal with values

*/

func main() {
	// 1. CREATING MAPS

	// Method 1: Using make (most common)
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	fmt.Println("Using make:", scores)

	// Method 2: Map literal
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	fmt.Println("Map literal:", ages)

	// Method 3: Empty map literal
	empty := map[string]int{}
	fmt.Println("Empty map:", empty, "is nil:", empty == nil) // false, it's not nil

	// DANGER: nil map - can read but NOT write
	var nilMap map[string]int
	fmt.Println("nil map:", nilMap, "is nil:", nilMap == nil) // true
	// nilMap["key"] = 1  // PANIC: assignment to entry in nil map

	// 2. BASIC OPERATIONS

	fmt.Println("\n--- Basic Operations ---")
	m := make(map[string]int)

	// Insert / Update
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println("After inserts:", m)

	// Update existing key
	m["one"] = 100
	fmt.Println("After update:", m)

	// Read
	fmt.Println("m[\"two\"] =", m["two"])

	// Read non-existent key returns zero value
	fmt.Println("m[\"missing\"] =", m["missing"]) // 0

	// Delete
	delete(m, "two")
	fmt.Println("After delete:", m)

	// Length
	fmt.Println("Length:", len(m))

	// 3. CHECKING IF KEY EXISTS (comma-ok idiom)

	fmt.Println("\n--- Checking key existence ---")
	inventory := map[string]int{
		"apples":  5,
		"bananas": 0, // exists but value is 0
	}

	// Problem: can't distinguish "key doesn't exist" from "value is zero"
	fmt.Println("oranges:", inventory["oranges"])  // 0
	fmt.Println("bananas:", inventory["bananas"])  // 0

	// Solution: comma-ok idiom
	val, exists := inventory["oranges"]
	fmt.Printf("oranges: value=%d, exists=%v\n", val, exists)

	val, exists = inventory["bananas"]
	fmt.Printf("bananas: value=%d, exists=%v\n", val, exists)

	// Common pattern: check before using
	if count, ok := inventory["apples"]; ok {
		fmt.Println("We have", count, "apples")
	} else {
		fmt.Println("No apples in inventory")
	}

	// 4. ITERATING OVER MAPS

	fmt.Println("\n--- Iteration ---")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	// Key and value
	fmt.Println("Key-value pairs:")
	for key, value := range colors {
		fmt.Printf("  %s: %s\n", key, value)
	}

	// Keys only
	fmt.Print("Keys: ")
	for key := range colors {
		fmt.Print(key, " ")
	}
	fmt.Println()

	// Values only
	fmt.Print("Values: ")
	for _, value := range colors {
		fmt.Print(value, " ")
	}
	fmt.Println()

	// 5. MAPS WITH STRUCT VALUES

	fmt.Println("\n--- Maps with structs ---")
	type Location struct {
		Lat, Long float64
	}

	places := map[string]Location{
		"NYC":    {40.7128, -74.0060},
		"London": {51.5074, -0.1278},
		"Tokyo":  {35.6762, 139.6503},
	}

	for city, loc := range places {
		fmt.Printf("%s: %.4f, %.4f\n", city, loc.Lat, loc.Long)
	}

	// 6. MAPS AS SETS

	fmt.Println("\n--- Map as Set ---")
	// Go doesn't have a built-in set type, use map[T]bool or map[T]struct{}

	// Using map[string]bool
	visited := map[string]bool{
		"Paris":  true,
		"Tokyo":  true,
		"London": true,
	}

	city := "Paris"
	if visited[city] {
		fmt.Println(city, "has been visited")
	}

	// Using map[string]struct{} (more memory efficient)
	seen := make(map[string]struct{})
	seen["apple"] = struct{}{}
	seen["banana"] = struct{}{}

	if _, ok := seen["apple"]; ok {
		fmt.Println("apple is in the set")
	}

	// 7. NESTED MAPS

	fmt.Println("\n--- Nested maps ---")
	// Map of maps
	users := map[string]map[string]string{
		"user1": {
			"name":  "Alice",
			"email": "alice@example.com",
		},
		"user2": {
			"name":  "Bob",
			"email": "bob@example.com",
		},
	}

	fmt.Println("user1 name:", users["user1"]["name"])

	// Adding to nested map - must initialize inner map first!
	users["user3"] = make(map[string]string)
	users["user3"]["name"] = "Charlie"
	users["user3"]["email"] = "charlie@example.com"

	for id, info := range users {
		fmt.Printf("%s: %s (%s)\n", id, info["name"], info["email"])
	}

	// 8. COUNTING WITH MAPS

	fmt.Println("\n--- Word counting ---")
	text := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}

	wordCount := make(map[string]int)
	for _, word := range text {
		wordCount[word]++ // zero value (0) + 1 works perfectly
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

	// 9. GROUPING WITH MAPS

	fmt.Println("\n--- Grouping ---")
	type Person struct {
		Name string
		City string
	}

	people := []Person{
		{"Alice", "NYC"},
		{"Bob", "London"},
		{"Charlie", "NYC"},
		{"Diana", "London"},
	}

	byCity := make(map[string][]string)
	for _, p := range people {
		byCity[p.City] = append(byCity[p.City], p.Name)
	}

	for city, names := range byCity {
		fmt.Printf("%s: %v\n", city, names)
	}

	// 10. CLEARING A MAP

	fmt.Println("\n--- Clearing a map ---")
	toClear := map[string]int{"a": 1, "b": 2}
	fmt.Println("Before clear:", toClear)

	// Method 1: Delete each key
	for k := range toClear {
		delete(toClear, k)
	}
	fmt.Println("After delete loop:", toClear)

	// Method 2: Reassign (previous map becomes garbage collected)
	toClear = make(map[string]int)
	fmt.Println("After reassign:", toClear)
}
