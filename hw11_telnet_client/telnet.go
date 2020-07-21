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

type client struct {
	connection net.Conn
	inScanner  *bufio.Scanner
	outScanner *bufio.Scanner
	in         io.ReadCloser
	out        io.Writer
	address    string
	timeout    time.Duration
}

func (c client) Close() error {
	if c.connection != nil {
		err := c.connection.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c client) Send() error {
	success := c.inScanner.Scan()
	// TODO: put on one line
	if success == false {
		c.Close()
	}

	message := c.inScanner.Text() + "\n"
	_, err := c.connection.Write([]byte(message))

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (c client) Receive() error {
	c.outScanner.Scan()
	// TODO: put on one line
	_, err := c.out.Write(c.outScanner.Bytes())
	c.out.Write([]byte("\n"))
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (c *client) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		fmt.Println("No connection!")
		return err
	}
	c.connection = conn

	inScanner := bufio.NewScanner(c.in)
	outScanner := bufio.NewScanner(c.connection)

	inScanner.Split(bufio.ScanLines)
	outScanner.Split(bufio.ScanLines)

	c.inScanner = inScanner
	c.outScanner = outScanner

	return nil
}

func createTelnetClient(in io.ReadCloser, out io.Writer, address string, timeout time.Duration) TelnetClient {
	return &client{nil, nil, nil, in, out, address, timeout}
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here

	//for x := range time.Tick(time.Second) {
	//	connection.Write([]byte(fmt.Sprintf("Time: %d", x)))
	//}

	//go func(conn net.Conn, in io.ReadCloser) {
	//	scanner := bufio.NewScanner(in)
	//	scanner.Split(bufio.ScanLines)
	//	for scanner.Scan() {
	//		// TODO: put on one line
	//		_, err = conn.Write(scanner.Bytes())
	//		conn.Write([]byte("\n"))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//	conn.Close()
	//	fmt.Println("Connection is closed due to EOF!")
	//}(conn, in)
	//
	//go func(conn net.Conn, out io.Writer) {
	//	scanner := bufio.NewScanner(conn)
	//	scanner.Split(bufio.ScanLines)
	//	for scanner.Scan() {
	//		// TODO: put on one line
	//		_, err = out.Write(scanner.Bytes())
	//		out.Write([]byte("\n"))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//}(conn, out)

	return createTelnetClient(in, out, address, timeout)
}

// Place your code here
// P.S. Author's solution takes no more than 50 lines
