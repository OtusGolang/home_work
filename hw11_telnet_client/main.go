package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Place your code here
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	//in := &bytes.Buffer{}
	//out := &bytes.Buffer{}
	var wg sync.WaitGroup
	wg.Add(1)

	tc := NewTelnetClient("0.0.0.0:4242", 10 * time.Second, os.Stdin, os.Stdout)

	go func(wg *sync.WaitGroup, tc TelnetClient) {
		defer wg.Done()
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		tc.Close()
	}(&wg, tc)

	wg.Wait()
}
