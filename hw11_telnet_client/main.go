package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Place your code here
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	//in := &bytes.Buffer{}
	//out := &bytes.Buffer{}
	var wg sync.WaitGroup
	wg.Add(1)

	NewTelnetClient("0.0.0.0:4242", 1, os.Stdin, os.Stdout)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	}(&wg)

	wg.Wait()
}
