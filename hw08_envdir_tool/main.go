package main

import (
	"log"
	"os"
)

const minArgs = 3

func main() {
	if len(os.Args) < minArgs {
		log.Fatalln("Необходимо минимум 3 аргумента. Использование: ./go-envdir /path/to/env/dir command arg1 arg2")
	}

	eDir := os.Args[1]
	args := os.Args[2:]
	env, err := ReadDir(eDir)
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(RunCmd(args, env))
}
