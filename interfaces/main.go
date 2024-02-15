package main

import "fmt"

/*

Interfaces
	An interface is a type that defines one or more method signatures.
	This is said to define a contract, but not the implementation.
	In an interface, methods (functions) are not implemented,
	just their signature is declared. A signature is the function name
	and parameter type(s).


	the ability to be used the interface is based solely on the method signature.
	You do not need to declare that a struct or type is going to implement
	an interface explicitly

*/

type Session interface {
	ChargeStart()
	ChargeEnd()
}

type EvChargeSession struct {
	Id    string
	Watts int
}

type HybridChargeSession struct {
	Id    string
	Watts int
	Hp    int
}

func (p EvChargeSession) ChargeStart() { // this implements the ChargeStart()
	fmt.Println("charging session initiated for ev", p.Id)
}

func (p EvChargeSession) ChargeEnd() { // this implements the ChargeEnd()
	fmt.Println("charging session ended for ev", p.Id)
}

func (p HybridChargeSession) ChargeStart() { // this implements the ChargeStart()
	fmt.Println("charging session initiated for hybrid", p.Id)
}

func (p HybridChargeSession) ChargeEnd() { // this implements the ChargeEnd()
	fmt.Println("charging session ended for hybrid", p.Id)
}

func ChargeInitiate(cs Session) { // this takes type of our interface
	cs.ChargeStart()
}

func ChargeTerminate(cs Session) { // this takes type of our interface
	cs.ChargeEnd()
}

func main() {
	cs := EvChargeSession{Id: "EV-420", Watts: 420}
	ChargeInitiate(cs)

	csh := HybridChargeSession{Id: "Hybrid-480", Watts: 420, Hp: 480}
	ChargeInitiate(csh)

	ChargeTerminate(cs)
	ChargeTerminate(csh)
}
