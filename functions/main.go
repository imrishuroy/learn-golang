package main

import "fmt"

// Higher order functions
/*
	The defination of Higher-order function is a function which does at least one of the following
	1. takes one or more functions as arguments
	2. returns a funcitons as its result

*/

type chargeSession struct {
	id    string
	watts int
	vin   string
}

func filter(csIn []chargeSession, f func(chargeSession) bool) []chargeSession {
	var csOut []chargeSession
	for _, v := range csIn {
		if f(v) == true {
			csOut = append(csOut, v)
		}
	}
	return csOut
}

func main() {

	// a real world example
	// In the function below, the second parameter to filter is a function which takes a chargeSession
	// as a parameter and returns a bool. This function determines whether a particular chargingSession
	// matches a criteria or not of being above a certain value (greater than 400).
	// We iterate through the chargingSession slice and pass each chargingSession as a
	// parameter to the function f. If this returns true, it means that that the
	// chargeSession has passed the filter criteria and he is added to the slice csOut.

	testExample()

	// Closures
	// Go also support anonymous fuctions, which can form closures.
	// Anonymous functions are useful when you want to define a function inline
	// without having to name it

	firstInt := intFunc()

	fmt.Println(firstInt())
	fmt.Println(firstInt())

	secondInt := intFunc()
	fmt.Println(secondInt())
	fmt.Println(secondInt())

	// First class functions

	charge(levelThree, "Tesla") // function as an argument
	charge(levelTwo, "F-150")

	a2 := func() {
		fmt.Println("Go does support first class functions")
	}

	a2()

	fmt.Printf("%T\n", a2)

	multiple(2, 3)
	multiple(3, 4, 5)

	nums := []int{6, 7, 8, 9}
	multiple(nums...)

	a, b := multipleVals() // assing both return values in one statement
	// We will see this quite often is when the first value returned is the
	// Result of the function, the second valur is an Error if there was one.
	fmt.Println("First value", a)
	fmt.Println("Second value", b)

	_, c := multipleVals() // here we ignore the first one
	fmt.Println("Just second", c)

}

// closer func

// It will print the same output twice. The function
// intFunc returns an anonymous function that returns a value multiplied by 2.
// It is a closure, because it closes over the variable i which is declared outside the
// anonymous function
func intFunc() func() int {
	i := 0
	return func() int {
		i++
		return i * 2
	}
}

func testExample() bool {
	cs1 := chargeSession{
		id:    "11111",
		watts: 420,
		vin:   "4Y1SL65848Z411439",
	}
	cs2 := chargeSession{
		id:    "22222",
		watts: 390,
		vin:   "3Y2RM78848Z411439",
	}
	cs3 := chargeSession{
		id:    "33333",
		watts: 401,
		vin:   "2Y2RM78848Z411439",
	}

	s := []chargeSession{cs1, cs2, cs3}
	f := filter(s, func(s chargeSession) bool {
		return s.watts > 400
	})
	fmt.Println(f)

	return false

}

func charge(f func() string, vehicleName string) {
	fmt.Println("I will charge my", vehicleName, f())
}

func levelThree() string {
	return "with a Level 3 charger"
}

func levelTwo() string {
	return "with a Level 2 charger"
}

// variadic implementation for multiplying an unknown numbers of parameters.
func multiple(nums ...int) {
	fmt.Println("Parameters", nums, " ")
	total := 1

	for _, num := range nums {
		total *= num
	}

	fmt.Println(" with total when multiplied", total)
}

func multipleVals() (int, int) {
	return 4, 9
}
