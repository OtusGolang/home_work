package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Init(path string) error
}

type Instance struct {
}

func (i *Instance) Init(path string) error {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)

		return err
	}

	log.SetOutput(file)
	fmt.Println("logger initialized, log file: ", path)

	return nil
}
