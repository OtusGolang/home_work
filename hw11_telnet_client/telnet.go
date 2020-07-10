package main

import "net"

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

//func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
func NewTelnetClient() TelnetClient {
	// Place your code here
	conn, _ := net.Dial("tcp", "0.0.0.0:4242")
	conn.Write([]byte("Hello Yan"))
	conn.Close()
	return nil
}

// Place your code here
// P.S. Author's solution takes no more than 50 lines
