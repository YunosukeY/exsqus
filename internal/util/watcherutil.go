package util

import (
	"os"

	"github.com/fsnotify/fsnotify"
)

func GetWatcher() (*fsnotify.Watcher, error) {
	path := os.Getenv("LOG_FILE_PATH")
	if path == "" {

	}

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
