package main

import "fmt"

// struct
// A struct is a type that contains one or more variables.
// It’s like a collection of variables. They are referred to as fields.

// While Go is not, on purpose, an object oriented language it does have structures,
// which can be associated with methods.

// When declaring the struct, we used uppercase names for the fields,
// otherwise those will be private to the package and not accessible to the outside.

type ChargeSessionDetails struct {
	StartTime string
	EndTime   string
}

type ChargeSession struct {
	Id    string
	Watts int
	csd   ChargeSessionDetails
}

func main() {
	cs := ChargeSession{Id: "123", Watts: 1000,
		csd: ChargeSessionDetails{
			StartTime: "2020-01-01",
			EndTime:   "2020-01-02"}}
	fmt.Println(cs)

	// Note you can declare an struct as empty by not providing any
	// values and then assign them later.
	cs2 := ChargeSession{}
	cs2.Id = "456"
	cs2.Watts = 2000
	fmt.Println(cs2)

	cs.IncreaseWatts(10)
	fmt.Println(cs)
	cs.IncreaseWatts(20)
	fmt.Println(cs)
	cs.NotifyEndSession()

	// Using the constructor
	cs3 := NewChargeSession("789", 3000)
	fmt.Println(cs3)
	cs3.IncreaseWatts(30)
	fmt.Println(cs3.Watts)

	// This returns a pointer to a new ChargeSession struct and assigns it to cs4.
	cs4 := new(ChargeSession)
	cs4.Id = "101112"
	cs4.Watts = 4000
	cs4.csd.StartTime = "2020-01-03"
	cs4.csd.EndTime = "2020-01-04"
	fmt.Println(*cs4)

}

// Receivers
// A receiver is a function that has a reference to the instance of the struct.
// Value Receiver
func (c ChargeSession) NotifyEndSession() {
	fmt.Println("Sending notification to id", c.Id, "end of session for total watts", c.Watts)
}

// Pointer Receiver
func (c *ChargeSession) IncreaseWatts(w int) {
	c.Watts += w
}

// a constructor function. This is a function that returns an instance of the desired type.
// Go does not have constructors as part of the language, but it does have
// the ability to create a function that returns a new value.
func NewChargeSession(id string, watts int) *ChargeSession {
	cs := ChargeSession{Id: id, Watts: watts}
	return &cs
}
