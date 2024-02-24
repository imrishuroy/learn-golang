package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

// You might have some setup of test data or cleanup required.
// That is done as follows in this example by adding a function called TestMain.
// As you see it receives a parameter of m *testing.M and when we call m.Run()
// the tests actually run. Before that data or other setup can be done.
// After that any post test cleanup can be implemented.
func TestMain(m *testing.M) {
	fmt.Println("Do your setup for your tests now")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Do your cleanup for your tests now", testTime)
	os.Exit(exitVal)
}

// tests use the func Testxxxx(t *Testing.t),
// format, where xxxx is a unique name for the test and starts
// with an uppercase letter, TestDivideZero.
func TestDivide(t *testing.T) {
	type testCases struct {
		nom  int
		den  int
		want int
	}
	for _, tc := range []testCases{
		{nom: 10, den: 5, want: 2},
		{nom: -12, den: 4, want: -3},
		{nom: 19, den: 2, want: 9},
		{nom: 2, den: 1, want: 2},
		{nom: 0, den: 2, want: 0},
	} {
		expected := tc.want
		result, _ := divide(tc.nom, tc.den)
		if result != expected {
			// While t.Error or its sibling t.Errorf mark a test as failed,
			// the test function continues running
			// This is usually the desired behavior. If you do think a test function
			// should stop processing as soon as a failure is found,
			// use the t.Fatal or t.Fatalf methods.
			t.Errorf("Expected %d but got %d", expected, result)
		}
	}
}

func TestDivideByZero(t *testing.T) {
	result, e := divide(10, 0)
	if e == nil {
		t.Errorf("Expected 0 but got %d", result)
	}
}
