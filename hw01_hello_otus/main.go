package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	reversedStr := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(reversedStr)
}
