package main

import (
	"errors"
	"fmt"
)

// rather than throwing exceptions, errors are treated as values

// error is an interface and anything that implements this interface is considered an error.

// There are two types of errors in Go.
// Anticipated errors which we have been looking at and unanticipated errors, called Panics.
var Error_Invalid_Denominator = errors.New("zero denominator")

type ErrorDivision struct {
	Num int
	Den int
	Msg string
}

// implement the error interface
func (e *ErrorDivision) Error() string {
	return e.Msg
}

// func divide(n, d int) (int, error) {
// 	if d == 0 {
// 		// return 0, errors.New("denominator cannot be zero")
// 		// return 0, fmt.Errorf("invalid denominator: %d", d)
// 		// return 0, Error_Invalid_Denominator

// 		return 0, &ErrorDivision{Num: n, Den: d, Msg: "invalid denominator"}
// 	}

// 	return n / d, nil
// }

// The deferred function calls recover and will result in a normal error message.
// This would normally result in a panic exit, but we can continue processing and
// if we had other logic in our program this would be executed.
func divide(n, d int) (int, error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic occurred we recovered:", e)
		}
	}()
	if d == 0 {
		panic(errors.New("divide by zero"))
	}
	return n / d, nil
}

func main() {
	// panic

	r, e := divide(1, 0)
	if e != nil {
		fmt.Println(e)
		panic(e)
	} else {
		fmt.Println("Remainder:", r)
	}

	// This will print this, seeing the deferred function executed before the exit.
	// defer func() {
	// 	fmt.Println("Let's cleanup program resources before going down")
	// }()

	// panic(errors.New("a bad error"))

	// error handling
	// r, e := divide(10, 0)
	// if e != nil {
	// 	var errDiv *ErrorDivision
	// 	switch {
	// 	case errors.As(e, &errDiv):
	// 		fmt.Printf("%d / %d is not a valid division statement: %s\n",
	// 			errDiv.Num, errDiv.Den, errDiv.Error())
	// 	default:
	// 		fmt.Printf("unexpected division error: %s\n", e)
	// 	}
	// 	return
	// }
	// fmt.Println("Remainder:", r)

	// r, e := divide(10, 0)
	// if e != nil {
	// 	switch {
	// 	case errors.Is(e, Error_Invalid_Denominator):
	// 		fmt.Println("divide by zero Error")
	// 	default:
	// 		fmt.Printf("unexpected division error: %s\n", e)
	// 	}
	// 	return
	// }

	// fmt.Println("Remainder:", r)

	// r, e := divide(10, 0)
	// if e != nil {
	// 	fmt.Println("Error:", e)
	// } else {
	// 	fmt.Println("Remainder:", r)
	// }
}
