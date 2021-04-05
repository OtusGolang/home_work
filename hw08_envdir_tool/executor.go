package main

import (
	"errors"
	"os"
	"os/exec"
)

const (
	errorCode   = 1
	successCode = 0
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return errorCode
	}

	c := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	for param, value := range env {
		var err error

		if value.NeedRemove {
			err = os.Unsetenv(param)
		} else {
			err = os.Setenv(param, value.Value)
		}

		if err != nil {
			return errorCode
		}
	}

	c.Env = os.Environ()
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin

	if err := c.Run(); err != nil {
		var exitErr *exec.ExitError

		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
		return errorCode
	}

	return successCode
}
