package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	err := Copy("testdata/input.txt", "out.txt", 0 ,0)

	require.Nil(t, err)
}

func TestNegativeOffset(t *testing.T) {
	err := Copy("testdata/input.txt", "out.txt", -10, 0)

	require.Error(t, err)
}

func TestNegativeLimit(t *testing.T) {
	err := Copy("testdata/input.txt", "out.txt", 0, -1)

	require.Error(t, err)
}

func TestUnsupportedFile(t *testing.T) {
	err := Copy("/testdata/nofile.txt", "out.txt", 0, 0)

	require.Equal(t, true, errors.Is(err, ErrUnsupportedFile))
}

func TestOffsetExceedsFileSize(t *testing.T) {
	err := Copy("testdata/input.txt", "out.txt", 1000000, 0)

	require.Equal(t, true, errors.Is(err, ErrOffsetExceedsFileSize))
}