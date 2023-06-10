package util

import "github.com/stretchr/testify/assert"

func assertEqual[T any](t assert.TestingT, expected, actual T, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, actual, msgAndArgs...)
}
