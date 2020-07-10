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

type hello struct {
	net.Conn
}

func (c hello) Close() error {
	return c.Close()
}

func (c hello) Send() error {
	return c.Close()
}

func (c hello) Receive() error {
	return c.Close()
}

func (c hello) Connect() error {
	return c.Close()
}

func NopTelnetClient(c net.Conn) TelnetClient {
	return hello{c}
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here
	conn, err := net.Dial("tcp", address)
	if err != nil  {
		fmt.Println("No connection!")
		return nil
	}
	for x := range time.Tick(time.Second) {
		conn.Write([]byte(fmt.Sprintf("Time: %d", x)))
	}
	return NopTelnetClient(conn)
}

// Place your code here
// P.S. Author's solution takes no more than 50 lines
