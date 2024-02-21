package main

import "fmt"

// func Last[T any](items []T) T {
// 	return items[len(items)-1]
// }

// The first part of the function Print in the square brackets
// is what is called a constraint.

// Constraints are just interfaces in reality
// and any is the same as interface{}.
// func Print[T any](s []T) {
// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// 	fmt.Println()
// }

func Stringify[T fmt.Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

type Person struct {
	Name string
	Age  int
}

type Role struct {
	Name  string
	Years int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (r Role) String() string {
	return fmt.Sprintf("In the role %v for %v years", r.Name, r.Years)
}

func main() {

	s := Stringify[Person]([]Person{{"Brian Enochson", 42}, {"Tom Smith", 29}})
	fmt.Println(s)

	s2 := Stringify([]Person{{"Brian Keith Enochson", 42}})
	fmt.Println(s2)

	s3 := Stringify[Role]([]Role{{"Developer", 20}, {"Architect", 5}})
	fmt.Println(s3)

	// Print[int]([]int{10, 20, 30, 40, 50})

	// Print[string]([]string{"ford", "chevy", "toyota", "honda"})

	// Print[float32]([]float32{10.5, 20.5, 30.5, 40.5, 50.5})

	// intSlice := []int{10, 20, 30, 40, 50}
	// firstInd := Last[int](intSlice)

	// fmt.Println(firstInd)

	// strSlice := []string{"ford", "chevy", "toyota", "honda"}
	// firstStr := Last[string](strSlice)

	// fmt.Println(firstStr)
}
