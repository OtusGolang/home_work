package main

import (
	"fmt"

	"golang.org/x/text/unicode/bidi"
)

func main() {
	fmt.Println(bidi.ReverseString("Hello Otus"))
}
