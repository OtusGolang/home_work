package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	args := getArgs()
	ctx, cancel := context.WithCancel(context.Background())

	tc := NewTelnetClient(args.host + ":" + args.port, time.Duration(args.timeout) * time.Second, os.Stdin, os.Stdout)
	err := tc.Connect()
	if err != nil {
		log.Fatal(err)
	}

	go func(tc TelnetClient, ctx context.Context, cancel context.CancelFunc) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := tc.Receive()
				if err != nil {
					cancel()
					return
				}
			}

		}
	}(tc, ctx, cancel)

	go func(tc TelnetClient, ctx context.Context, cancel context.CancelFunc) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := tc.Send()
				if err != nil {
					cancel()
					return
				}
			}
		}
	}(tc, ctx, cancel)

	go handleSignals(tc, cancel)

	<-ctx.Done()
}

func handleSignals(tc TelnetClient, cancel context.CancelFunc) {
	defer cancel()
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	err := tc.Close()
	if err != nil {
		log.Fatal(err)
	}
}

type Args struct {
	host string
	port string
	timeout int
}

func getArgs() *Args {
	timeout := flag.String("timeout", "10s", "connection timeout")
	flag.Parse()
	otherArgs := flag.Args()

	args := Args{
		host: otherArgs[0],
		port: otherArgs[1],
		timeout: func() int {
			timeoutStr := *timeout
			timeoutInt, err := strconv.Atoi(timeoutStr[:len(timeoutStr)-1])
			if err != nil {
				log.Fatal("Timeout format is not correct")
			}
			return timeoutInt
		}(),
	}

	if args.host == "" {
		log.Fatal("Specify 1st parameter: the host")
	}
	if args.port == "" {
		log.Fatal("Specify 2nd parameter: the port")
	}

	return &args
}
