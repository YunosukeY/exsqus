package util

import (
	"bufio"
	"io"
)

func SkipAll(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
	}
}
