package main

import (
	"bookshelf/config"
	"bookshelf/server/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()

	appRouter := router.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Println("Starting server :8080")

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
