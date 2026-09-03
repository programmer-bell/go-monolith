package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/programmer-bell/go-monolith/internal/config"
	"github.com/programmer-bell/go-monolith/internal/db"
	"github.com/programmer-bell/go-monolith/internal/handlers"
	"github.com/programmer-bell/go-monolith/internal/middleware"
)

func main() {

	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect:%v", err)
	}

	loghandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	logger := slog.New(loghandler)
	slog.SetDefault(logger)

	fmt.Println("Database connected.")
	fmt.Println("Starting olx server...")

	lh := handlers.NewListingHandler(db, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.HealthzHandler)
	mux.HandleFunc("GET /listing", lh.List)
	mux.HandleFunc("DELETE /listings/{id}", lh.Delete)

	handler := middleware.RequestId(mux)
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server is listenning %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
