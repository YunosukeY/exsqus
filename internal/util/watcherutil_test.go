package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWatcher(t *testing.T) {
	watcher, err := GetWatcher([]string{"watcherutil.go"})

	assert.Nil(t, err)
	assertEqual(t, []string{"watcherutil.go"}, watcher.WatchList())
}
