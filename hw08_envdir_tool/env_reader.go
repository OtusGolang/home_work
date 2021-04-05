package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := Environment{}

	for _, fInfo := range files {
		// Пропускаем файл с = или если это директория
		if fInfo.IsDir() || strings.Contains(fInfo.Name(), "=") {
			continue
		}

		v, err := getVal(dir, fInfo)
		if err != nil {
			return nil, err
		}

		ev := EnvValue{}
		ev.Value = v

		if v == "" {
			ev.NeedRemove = true
		}

		env[fInfo.Name()] = ev
	}

	return env, nil
}

func getVal(dir string, fInfo os.FileInfo) (string, error) {
	if fInfo.Size() == 0 {
		return "", nil
	}

	fPath := filepath.Join(dir, fInfo.Name())
	file, err := os.Open(fPath)
	// Доступен ли на чтение
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	val := string(scanner.Bytes())
	val = strings.ReplaceAll(val, "\x00", "\n")
	val = strings.TrimRight(val, " \t")

	return val, nil
}
