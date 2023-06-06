package app

import (
	"log"
	"os"

	"github.com/YunosukeY/explain-slow-query/internal/util"
	"github.com/fsnotify/fsnotify"
)

func Run() {
	path := os.Getenv("LOG_FILE_PATH")
	if path == "" {

	}
	log.Println("log file path:", path)

	db, err := util.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("DB connected")

	watcher, err := util.GetWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	log.Println("Start watching")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write && event.Name == path {
				l, err := util.GetLastQueryLog()
				if err != nil {
					log.Println("Failed to get last query log:", err)
					continue
				}
				log.Println("Log:", l)
				plan, err := util.GetPlan(db, l.Query)
				if err != nil {
					log.Println("Failed to get plan:", err)
					continue
				}
				log.Println("Plan:", plan)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}
