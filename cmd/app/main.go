package main

import (
	"bookshelf/config"
	"bookshelf/server/app"
	"bookshelf/server/router"
	lr "bookshelf/util/logger"
	"fmt"
	"net/http"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Server.Debug)

	application := app.New(logger)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
