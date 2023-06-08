package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWatcher(t *testing.T) {
	watcher, err := GetWatcher("watcherutil.go")

	assert.Nil(t, err)
	assert.Equal(t, []string{"watcherutil.go"}, watcher.WatchList())
}
