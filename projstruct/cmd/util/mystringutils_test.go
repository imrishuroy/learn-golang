package util

import "testing"

func TestReverseString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello Go Project Structure", "erutcurtS tcejorP oG olleH"},
		{"Hello, Brian", "nairB ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseString(c.in)
		if got != c.want {
			t.Errorf("ReverseString(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestUpperCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello Go Project Structure", "HELLO GO PROJECT STRUCTURE"},
		{"Hello, Brian", "HELLO, BRIAN"},
		{"", ""},
	}
	for _, c := range cases {
		got := UpperCase(c.in)
		if got != c.want {
			t.Errorf("UpperCase(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
