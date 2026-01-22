package main

import (
	"context"
	"culand-internship/internal/core"
	"culand-internship/internal/core/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := core.Start(cfg)
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}

	serverErr := make(chan error, 1)

	go func() {
		log.Printf("server started on :%d", cfg.Port)
		serverErr <- app.Server.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		log.Printf("server error: %v", err)

	case <-stop:
		log.Println("shutdown signal received")
	}

	ctx := context.Background()
	if err := app.Stop(ctx); err != nil {
		log.Printf("failed to stop app: %v", err)
	}

	log.Println("application stopped")
}
