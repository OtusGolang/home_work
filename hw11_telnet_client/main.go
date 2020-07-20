package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

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

func main() {
	args := getArgs()
	var wg sync.WaitGroup
	wg.Add(1)

	tc := NewTelnetClient(args.host + ":" + args.port, time.Duration(args.timeout) * time.Second, os.Stdin, os.Stdout)
	tc.Connect()
	//fmt.Printf("%+v\n", tc)
	go func(wg *sync.WaitGroup, tc TelnetClient) {
		defer wg.Done()
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		tc.Close()
	}(&wg, tc)

	wg.Wait()
}
