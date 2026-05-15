package main

import "fmt"

/*
  GENERICS IN GO (Go 1.18+)

KEY POINTS:
  - Generics allow writing functions/types that work with multiple types
  - Type parameters are specified in square brackets [T]
  - Constraints define what operations are allowed on type parameters
  - Reduces code duplication while maintaining type safety

SYNTAX:
  func Name[T constraint](param T) T { }     // generic function
  type Name[T constraint] struct { }          // generic type

COMMON CONSTRAINTS:
  - any              : any type (alias for interface{})
  - comparable       : types supporting == and !=
  - Ordered : types supporting < > <= >=

WHY GENERICS:
  Before generics, you had to either:
  1. Write separate functions for each type (duplication)
  2. Use interface{} and lose type safety
  3. Use code generation

*/

func main() {
	// 1. THE PROBLEM: CODE DUPLICATION

	fmt.Println("--- The Problem ---")

	// Without generics, we need separate functions for each type
	intSlice := []int{10, 20, 15, -10}
	strSlice := []string{"Go", "is", "awesome"}

	fmt.Println("IndexInt:", IndexInt(intSlice, 15))
	fmt.Println("IndexString:", IndexString(strSlice, "is"))

	// 2. GENERIC FUNCTION - THE SOLUTION

	fmt.Println("\n--- Generic Index Function ---")

	// One function works for all comparable types
	fmt.Println("Index[int]:", Index(intSlice, 15))
	fmt.Println("Index[string]:", Index(strSlice, "is"))
	fmt.Println("Index (not found):", Index(intSlice, 999))

	// Type inference - compiler figures out the type
	fmt.Println("Index (inferred):", Index([]float64{1.1, 2.2, 3.3}, 2.2))

	// 3. MULTIPLE TYPE PARAMETERS

	fmt.Println("\n--- Multiple Type Parameters ---")

	m := map[string]int{"one": 1, "two": 2, "three": 3}

	keys := MapKeys(m)
	values := MapValues(m)

	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)

	// 4. CONSTRAINTS

	fmt.Println("\n--- Constraints ---")

	// 'any' constraint - accepts any type
	PrintAny(42)
	PrintAny("hello")
	PrintAny([]int{1, 2, 3})

	// 'comparable' constraint - types that support ==
	fmt.Println("Contains 2:", Contains([]int{1, 2, 3}, 2))
	fmt.Println("Contains x:", Contains([]string{"a", "b", "c"}, "x"))

	// Custom Number constraint
	fmt.Println("Sum ints:", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println("Sum floats:", Sum([]float64{1.1, 2.2, 3.3}))

	// 5. GENERIC MIN/MAX

	fmt.Println("\n--- Min/Max ---")

	fmt.Println("Min(3, 1, 4, 1, 5):", Min(3, 1, 4, 1, 5))
	fmt.Println("Max(3, 1, 4, 1, 5):", Max(3, 1, 4, 1, 5))
	fmt.Println("Min strings:", Min("banana", "apple", "cherry"))

	// 6. GENERIC TYPES (STRUCTS)

	fmt.Println("\n--- Generic Types ---")

	// Stack of integers
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println("Int stack:", intStack.items)
	fmt.Println("Pop:", intStack.Pop())
	fmt.Println("After pop:", intStack.items)

	// Stack of strings
	strStack := &Stack[string]{}
	strStack.Push("hello")
	strStack.Push("world")
	fmt.Println("String stack:", strStack.items)

	// 7. GENERIC PAIR TYPE

	fmt.Println("\n--- Generic Pair ---")

	pair1 := Pair[string, int]{First: "age", Second: 30}
	pair2 := Pair[string, string]{First: "name", Second: "Alice"}

	fmt.Println("Pair 1:", pair1.First, "=", pair1.Second)
	fmt.Println("Pair 2:", pair2.First, "=", pair2.Second)

	// Swap
	swapped := pair1.Swap()
	fmt.Printf("Swapped: %d = %s\n", swapped.First, swapped.Second)

	// 8. GENERIC MAP/FILTER/REDUCE

	fmt.Println("\n--- Map/Filter/Reduce ---")

	nums := []int{1, 2, 3, 4, 5}

	// Map: transform each element
	doubled := Map(nums, func(x int) int { return x * 2 })
	fmt.Println("Doubled:", doubled)

	// Map with type change
	asStrings := Map(nums, func(x int) string { return fmt.Sprintf("#%d", x) })
	fmt.Println("As strings:", asStrings)

	// Filter: keep elements matching predicate
	evens := Filter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println("Evens:", evens)

	// Reduce: combine all elements
	sum := Reduce(nums, 0, func(acc, x int) int { return acc + x })
	fmt.Println("Sum:", sum)

	product := Reduce(nums, 1, func(acc, x int) int { return acc * x })
	fmt.Println("Product:", product)

	// 9. TYPE SETS (INTERFACE CONSTRAINTS)

	fmt.Println("\n--- Type Sets ---")

	// Custom constraint using type sets
	fmt.Println("Double int:", Double(5))
	fmt.Println("Double float:", Double(3.14))
	fmt.Println("Double string:", Double("Go"))

	// 10. PRACTICAL: GENERIC CACHE

	fmt.Println("\n--- Generic Cache ---")

	cache := NewCache[string, int]()
	cache.Set("a", 1)
	cache.Set("b", 2)

	if val, ok := cache.Get("a"); ok {
		fmt.Println("cache[a] =", val)
	}

	if _, ok := cache.Get("c"); !ok {
		fmt.Println("cache[c] not found")
	}

	fmt.Println("All keys:", cache.Keys())
}

// NON-GENERIC FUNCTIONS (THE OLD WAY)

func IndexInt(s []int, x int) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func IndexString(s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// GENERIC FUNCTIONS

// Index returns the index of x in slice s, or -1 if not found
// [T comparable] means T must support == operator
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// MapKeys returns all keys from a map
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues returns all values from a map
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// PrintAny prints any value (any = interface{})
func PrintAny[T any](v T) {
	fmt.Printf("  PrintAny: %v (type: %T)\n", v, v)
}

// Contains checks if slice contains element
func Contains[T comparable](s []T, elem T) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

// Number constraint for numeric types
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Ordered constraint for types that support < > <= >=
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// Sum adds all numbers in a slice
func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// Min returns the minimum value
func Min[T Ordered](values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value
func Max[T Ordered](values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// GENERIC TYPES

// Stack is a generic LIFO data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Pair holds two values of potentially different types
type Pair[T, U any] struct {
	First  T
	Second U
}

func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{First: p.Second, Second: p.First}
}

// FUNCTIONAL UTILITIES

// Map transforms each element using a function
func Map[T, U any](s []T, fn func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

// Filter keeps elements that match the predicate
func Filter[T any](s []T, pred func(T) bool) []T {
	var result []T
	for _, v := range s {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce combines elements into a single value
func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range s {
		result = fn(result, v)
	}
	return result
}

// TYPE SETS

// Addable types that support + operator
type Addable interface {
	~int | ~float64 | ~string
}

// Double doubles a value (works with int, float64, string)
func Double[T Addable](v T) T {
	return v + v
}

// GENERIC CACHE

type Cache[K comparable, V any] struct {
	data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{data: make(map[K]V)}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache[K, V]) Keys() []K {
	return MapKeys(c.data)
}
