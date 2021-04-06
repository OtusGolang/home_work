package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const fTimeout = "timeout"

func main() {
	timeout := flag.Duration(fTimeout, 10*time.Second, "таймаут подключения к серверу")

	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("некорректные аргументы: укажите адрес и порт")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	addr := net.JoinHostPort(host, port)

	log.Printf("подключение к %s\n", addr)

	client := NewTelnetClient(addr, *timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		log.Fatalf("не удалось подключиться к %v, %v", addr, err)
	}

	log.Printf("успешно подключен")

	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatalf("не удалось закрыть соединение %v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	go listenStopSignal(cancel)
	go receive(client, cancel)
	go send(client, cancel)

	<-ctx.Done()
}

// Слушатель сигнала остановки программы.
func listenStopSignal(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	cancel()
}

func send(client TelnetClient, cancel context.CancelFunc) {
	if err := client.Send(); err != nil {
		log.Println(fmt.Errorf("не удалось отправить: %w", err))
	}
	cancel()
}

func receive(client TelnetClient, cancel context.CancelFunc) {
	if err := client.Receive(); err != nil {
		log.Println(fmt.Errorf("не удалось получить: %w", err))
	}
	cancel()
}
