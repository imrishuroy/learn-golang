package main

// We can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T.

type ChargeSession2 struct {
	Id    string
	watts int
	vin   string
}

func (cs *ChargeSession2) Charge(f int) {
	cs.watts = cs.watts + f
}
