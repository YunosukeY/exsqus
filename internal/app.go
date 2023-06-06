package app

import (
	"github.com/YunosukeY/explain-slow-query/internal/util"
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger()

	path := util.GetLogFilePath()
	log.Info().Str("log file path", path).Send()

	db, err := util.GetDB()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer db.Close()
	log.Info().Msg("MySQL connected")

	watcher, err := util.GetWatcher()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer watcher.Close()
	log.Info().Msg("Start watching")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write && event.Name == path {
				l, err := util.GetLastQueryLog()
				if err != nil {
					log.Info().Err(err).Msg("Failed to get last query log")
					continue
				}
				log.Info().Interface("log", l).Send()
				plan, err := util.GetPlan(db, l.Query)
				if err != nil {
					log.Info().Err(err).Msg("Failed to get plan")
					continue
				}
				log.Info().Interface("plan", plan).Send()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Info().Err(err).Msg("Watch error")
		}
	}
}
