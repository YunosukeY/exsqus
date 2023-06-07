package util

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLastQueryLog(t *testing.T) {
}

func TestGetTime(t *testing.T) {
	test := "# Time: 2023-06-07T11:58:58.688716Z\n"
	reader := bufio.NewReader(strings.NewReader(test))

	time, err := getTime(reader)
	assert.Nil(t, err)
	assert.Equal(t, "2023-06-07T11:58:58.688716Z", time)
}

func TestGetQueryTime(t *testing.T) {
}

func TestGetQuery(t *testing.T) {
}
