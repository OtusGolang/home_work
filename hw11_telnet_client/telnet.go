package main

import (
	"errors"
	"io"
	"time"
)

var ErrConnectionClosed = errors.New("connection closed by peer")

type TelnetClient interface {
	// Place your code here
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here
	return nil
}

// Place your code here
