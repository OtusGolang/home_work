package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("read", func(t *testing.T) {
		expected := Environment{
			"BAR":   {"bar", false},
			"UNSET": {"", true},
			"EMPTY": {"", true},
			"FOO":   {"   foo\nwith new line", false},
			"HELLO": {"\"hello\"", false},
		}
		env, err := ReadDir(
			"./testdata/env")
		require.NoError(t, err)
		require.Equal(t, expected, env)
	})

	t.Run("path not found", func(t *testing.T) {
		var expected Environment
		env, err := ReadDir("./testdata/no_file")
		require.Error(t, err)
		require.Equal(t, expected, env)
	})
}
