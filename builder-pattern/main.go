package main

import "fmt"

//  Builder Pattern
/*
	The Builder pattern is a powerful pattern that allows you to create complex objects
	with many optional parameters. What makes the builder pattern so effective is that
	you do not need to provide all parameters, you can provide defaults for certain fields
	and can implement validation at field or class level.

*/
type ChargeSession struct {
	ID    string
	Watts int
	Vin   string
}

type ChargeSessionBuilder struct {
	ChargeSession ChargeSession
}

// NewCsBuilder to create a new instance of that with an empty struct in it.
func NewCsBuilder() *ChargeSessionBuilder {
	return &ChargeSessionBuilder{
		ChargeSession: ChargeSession{},
	}
}

func (rb *ChargeSessionBuilder) SetId(id string) *ChargeSessionBuilder {
	// can perform id validation here if needed
	rb.ChargeSession.ID = id
	return rb
}

func (rb *ChargeSessionBuilder) SetWatts(watts int) *ChargeSessionBuilder {
	// can perform watts validation here if needed
	rb.ChargeSession.Watts = watts
	return rb
}

func (rb *ChargeSessionBuilder) SetVin(vin string) *ChargeSessionBuilder {
	// can perform vin validation here if needed
	rb.ChargeSession.Vin = vin
	return rb
}

// Finally the build() function ties it altogether. As a bonus, here we can check
// for missing values for optional fields and provide sensible defaults
func (rb *ChargeSessionBuilder) Build() ChargeSession {
	if rb.ChargeSession.Watts == 0 {
		rb.ChargeSession.Watts = -1
	}
	if len(rb.ChargeSession.Vin) == 0 {
		rb.ChargeSession.Vin = "Unknown"
	}
	return rb.ChargeSession
}

func main() {
	rb := NewCsBuilder()
	rb.SetId("123").SetWatts(420).SetVin("4Y1SL65848Z411439")
	fmt.Printf("charge session is %+v\n", rb.Build())

	rb = NewCsBuilder()
	rb.SetId("123")
	fmt.Printf("charge session is %+v\n", rb.Build())
}
