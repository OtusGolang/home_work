package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	var s = "Hello, OTUS!"
	fmt.Println(reverse.String(s))
}
