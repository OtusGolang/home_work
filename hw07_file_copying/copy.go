package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

const (
	bufsize = 1 << 10
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if offset < 0 {
		return fmt.Errorf("offset is less than zero")
	}
	if limit < 0 {
		return fmt.Errorf("limit is less than zero")
	}

	from, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer from.Close()

	to, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer to.Close()

	// Параметры файла
	fs, _ := from.Stat()

	if fs.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit+offset > fs.Size() {
		limit = fs.Size() - offset
	}

	o, err := from.Seek(offset, 0)
	if err != nil || o != offset {
		return err
	}

	fmt.Printf("Coping file %v to file %v", fromPath, toPath)
	fmt.Println()

	for i := offset; i < offset+limit; i += bufsize {
		copylen := int64(math.Min(bufsize, float64(limit)))
		c, err := io.CopyN(to, from, copylen)
		if (err != nil && err != io.EOF) || c > limit {
			return err
		}
	}
	return nil
}
