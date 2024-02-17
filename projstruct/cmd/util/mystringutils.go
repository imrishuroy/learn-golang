package util

import (
	strUtil "github.com/agrison/go-commons-lang/stringUtils"
)

// Exported identifiers with initial capital letter are visible
// outside of the package, and can be accessed from other packages.
func ReverseString(s string) string {
	return strUtil.Reverse(s)
}

func UpperCase(s string) string {
	return strUtil.UpperCase(s)
}
