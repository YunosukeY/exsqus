package util

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func RunCommonHandler() {
	http.HandleFunc("/healthz", healthCheck)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
}
