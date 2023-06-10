package util

import (
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
	reader := strings.NewReader(test)

	SkipAll(reader)

	bs, err := io.ReadAll(reader)
	assert.Nil(t, err)
	assertEqual(t, []byte{}, bs)
}
