package util

import (
	"github.com/fsnotify/fsnotify"
)

func GetWatcher(path string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	err = watcher.Add(path)
	if err != nil {
		watcher.Close()
		return nil, err
	}

	return watcher, nil
}
