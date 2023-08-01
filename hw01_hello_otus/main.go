package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	str := "Hello, OTUS!"
	revstring := reverse.String(str)
	fmt.Println(revstring)
}
