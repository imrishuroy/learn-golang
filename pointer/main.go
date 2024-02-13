package main

import "fmt"

type ChargeSession struct {
	Id    string
	Watts int
}

func ModifySession(cs *ChargeSession) {
	cs.Watts += 10
}

func main() {

	// ------- check pointer nil
	var prt2 *int

	if prt2 == nil {
		fmt.Println("Pointer is nill")
	} else {
		fmt.Println("Pointer in not nill")
	}

	// Note: We need to be careful of not using a pointer
	// unitl is has been assiged an address

	// ------- Poiner receiver
	cs2 := ChargeSession2{"1212", 410, "UNKNOWN"}
	cs2.Charge(10)
	fmt.Println(cs2.watts)

	// -------- A common use for pointer when working with large data structure like
	// struct which can have many fields.
	cs := ChargeSession{"11111", 420}
	s := &cs
	fmt.Println("ID", (*s).Id)
	fmt.Println("Watts:", s.Watts) // other more common way

	ModifySession(s)
	fmt.Println("After modify, Watts:", s.Watts)

	// allocate a ChargingSession struct using new and assign the address to the pointer t
	t := new(ChargeSession)
	t.Id = "22222"
	t.Watts = 389
	fmt.Println("ID:", t.Id)
	fmt.Println("Watts", t.Watts)

	// -------- Go Provide the new (T) function that can be used to create a pointer
	// to a newly allocated zero value of a specified type

	var ptr = new(int)
	fmt.Println("valur of 'ptr':", *ptr) // default is zero
	*ptr = 10
	fmt.Println("value of 'ptr':", *ptr)

	// -------- One use for pointer is passing them as arguments to functions
	// By passing just the pointer, it allows the function to modify the orignal data directly.

	num := 15
	modifyValue(&num)
	fmt.Println("modified value is", num)

	// ------- Pointer Example -------- //
	// taking a normal variable
	x := 123

	// initilizing a pointer
	p := &x // & is know as address operator

	// display the result
	fmt.Println("Value stored in x is :", x)
	fmt.Println("Address of x in ", &x)
	fmt.Println("Value stored in variable p is ", p) // the value is an address
	fmt.Println("Value p points to (*p) is ", *p)
}

func modifyValue(ptr *int) {
	*ptr = *ptr * 3
}
