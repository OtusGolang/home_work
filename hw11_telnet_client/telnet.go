package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type client struct {
	addr    string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &client{
		addr:    address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *client) Connect() error {
	var err error

	c.conn, err = net.DialTimeout("tcp", c.addr, c.timeout)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Send() error {
	r := bufio.NewReader(c.in)
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("EOF")
				return nil
			}
			return err
		}

		_, err = c.conn.Write([]byte(str))
		if err != nil {
			return err
		}
	}
}

func (c *client) Receive() error {
	r := bufio.NewReader(c.conn)
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("EOF")
				return nil
			}
			return err
		}

		_, err = c.out.Write([]byte(str))
		if err != nil {
			return err
		}
	}
}
