package util

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipAll(t *testing.T) {
	test := `1
	2
	3
	4
	5`
	reader := bufio.NewReader(strings.NewReader(test))

	err := SkipAll(reader)
	assert.Nil(t, err)

	_, err = reader.ReadString('\n')
	assert.Equal(t, io.EOF, err)
}
