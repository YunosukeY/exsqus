package util

import (
	"bufio"
	"io"
)

func SkipAll(reader *bufio.Reader) error {
	for {
		_, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}
