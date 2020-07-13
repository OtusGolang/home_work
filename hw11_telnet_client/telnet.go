package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	//for x := range time.Tick(time.Second) {
	//	conn.Write([]byte(fmt.Sprintf("Time: %d", x)))
	//}

	go func(conn net.Conn, in io.ReadCloser) {
		scanner := bufio.NewScanner(in)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			// TODO: put on one line
			_, err = conn.Write(scanner.Bytes())
			conn.Write([]byte("\n"))
			if err != nil {
				log.Fatal(err)
			}
		}
	}(conn, in)

	go func(conn net.Conn, out io.Writer) {
		scanner := bufio.NewScanner(conn)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			// TODO: put on one line
			_, err = out.Write(scanner.Bytes())
			out.Write([]byte("\n"))
			if err != nil {
				log.Fatal(err)
			}
		}
	}(conn, out)

	return NopTelnetClient(conn)
}

// Place your code here
// P.S. Author's solution takes no more than 50 lines
