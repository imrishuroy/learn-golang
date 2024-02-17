package main

import (
	strUtils "example/rishu/session/cmd/util"
	"fmt"
)

func main() {
	s := "Hello, Go Project Structure!"
	fmt.Println(s)
	s2 := strUtils.ReverseString(s)
	fmt.Println(s2)
	s3 := strUtils.UpperCase(s)
	fmt.Println(s3)
}
