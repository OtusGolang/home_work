package logger

import (
	"log"
	"os"
)

type Logger interface {
	Init(path string) error
}

type LoggerInstance struct {
}

func (i *LoggerInstance) Init(path string) error {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)

		return err
	}

	log.SetOutput(file)
	return nil
}
