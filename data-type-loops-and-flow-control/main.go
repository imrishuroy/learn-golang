// https://medium.com/gitconnected/learning-go-part-two-types-loops-and-flow-control-45251c7adff8
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

const Pi = 3.14159

func main() {

	// defer ( defer is used to ensure that a function call is performed later
	// in a program’s execution, usually for purposes of cleanup.)
	defer fmt.Println("This will be printed at the end")

	// Deferred function calls are pushed onto a stack. When a function returns,
	// its deferred calls are executed in last-in-first-out order.
	defer fmt.Println("This will be printed second")

	// flow control

	// 1. If Else
	// If else do not need parentheses, just like for loop
	// but blocks of code do need the curly braces, even if one statement.

	number := 10
	if number%2 == 0 {
		fmt.Println("Even Number")

	} else {
		fmt.Println("Odd Number")
	}

	limit := 10
	// we can also declare a variable in if condition
	if x := 4; x < limit {
		fmt.Println("x is less than limit")
	} else {
		fmt.Println("x is greater than limit")

	}

	// 2.Switch Case
	// An important difference to many languages is that Go’s switch cases
	// need not be constants and the values need not be integers.
	// and there is no need of break statement

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	// Here is an example where the switch with a condition,
	// where it evaluated the variable os.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Other OS", runtime.GOOS)
	}

	// while loop
	// Eg - 1
	i := 0
	for i < 5 {
		fmt.Println("Value of i is: ", i)
		i++
	}

	// Eg - 2 ( use of break )
	j := 0
	for true {
		fmt.Println("Value of j is: ", j)
		j++
		if j > 10 {
			break
		}
	}

	// loops
	numbers := [5]int{4, 8, 12, 16, 20}

	// _ is used to ignore the index,
	// as Go won't compile when a variable is declared and not used
	for _, item := range numbers {
		fmt.Println("Value of item is: ", item)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is: ", i)
	}

	const name3 = "Rishu"
	fmt.Println("Hey,", name3, "the value if Pi is", Pi)

	// type inference
	var name = "Rishu"
	name2 := "Prince"
	fmt.Println("This is a type ", reflect.TypeOf(name), "with value", name)
	fmt.Println("This is a type ", reflect.TypeOf(name2), "with value", name2)
	// we can also do this with %T ( type ) &v ( value )
	fmt.Printf("This is a type %T with value %v\n", name, name)

	var x int = 45 / 10
	var y = 7 - x
	z := 32 / 8
	z = 7
	fmt.Println("Value of x, y, and z is: ", x, y, z)

	// assign multiple variables
	a, b, c := 5, 7, 9
	fmt.Println("Value of a, b, and c is: ", a, b, c)

	floatNum := 3.14
	fmt.Printf("Float number %f\n", floatNum)

	// The allowed specifiers for printf follow these conventions.

	// %v — formats the value in a default format
	// %d — formats decimal integers
	// %t — formats true or false values
	// %g — formats the floating-point numbers
	// %s — formats string values
	// %b — formats base 22 numbers
	// %o — formats base 88 numbers
}
