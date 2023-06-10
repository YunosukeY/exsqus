package app

import (
	"os"

	"github.com/YunosukeY/exsqus/internal/util"
	"github.com/fsnotify/fsnotify"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func Run() {
	util.SetupLogger()

	paths := util.GetLogFilePaths()
	log.Info().Interface("log file paths", paths).Send()

	c := util.GetConfig()
	db, err := util.GetDB(c)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer db.Close()
	log.Info().Msg("MySQL connected")

	watcher, err := util.GetWatcher(paths)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer watcher.Close()
	log.Info().Msg("Start watching")

	fmap := map[string](*os.File){}
	for _, p := range paths {
		file, err := os.Open(p)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		defer file.Close()
		util.SkipAll(file)

		fmap[p] = file
	}

	go util.RunCommonHandler()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				file, ok := fmap[event.Name]
				if !ok {
					continue
				}

				id, err := uuid.NewRandom()
				if err != nil {
					log.Err(err).Send()
				}

				l, err := util.GetLastQueryLog(file)
				if err != nil {
					log.Err(err).Msg("Failed to get last query log")
					continue
				}
				l.Id = id.String()
				log.Info().Interface("log", l).Send()

				plan, err := util.GetPlan(db, l.Query)
				if err != nil {
					log.Err(err).Msg("Failed to get plan")
					continue
				}
				plan.Id = id.String()
				log.Info().Interface("plan", plan).Send()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}

			log.Err(err).Msg("Watch error")
		}
	}
}
