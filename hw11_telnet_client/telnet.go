package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here
	conn, err := net.Dial("tcp", address)
	if err != nil  {
		fmt.Println("No connection!")
		return nil
	}
	conn.Write([]byte("Hello Yan"))
	conn.Close()
	return nil
}

// Place your code here
// P.S. Author's solution takes no more than 50 lines
