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
	if success == false {
		fmt.Println("The end!")
		err := c.Close()
		if err != nil {
			log.Fatal(err)
		}
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
	return createTelnetClient(in, out, address, timeout)
}
