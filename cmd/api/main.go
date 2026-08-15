package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/programmer-bell/go-monolith/internal/config"
	"github.com/programmer-bell/go-monolith/internal/handlers"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println("Starting olx server")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.HealthzHandler)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server is listenning %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
