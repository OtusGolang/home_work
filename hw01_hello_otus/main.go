package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	str := "Hello, OTUS!"
	str = stringutil.Reverse(str)
	fmt.Println(str)
}
