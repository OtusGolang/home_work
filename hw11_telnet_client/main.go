package main

import (
	"bytes"
	"io/ioutil"
)

func main() {
	// Place your code here
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	NewTelnetClient("0.0.0.0:4242", 1, ioutil.NopCloser(in), out)
}
