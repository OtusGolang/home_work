package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	msg := "Hello, OTUS!"
	msg = stringutil.Reverse(msg)

	fmt.Printf(msg)
}
