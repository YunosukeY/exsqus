package util

import (
	"github.com/fsnotify/fsnotify"
)

func GetWatcher(paths []string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	for _, p := range paths {
		err = watcher.Add(p)
		if err != nil {
			watcher.Close()
			return nil, err
		}
	}

	return watcher, nil
}
