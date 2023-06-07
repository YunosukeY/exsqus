package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWatcher(t *testing.T) {
	watcher, err := GetWatcher("./")

	assert.Nil(t, err)
	assert.Equal(t, []string{"./"}, watcher.WatchList())
}
