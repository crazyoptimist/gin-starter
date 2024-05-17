package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gin-starter/internal/config"
	"gin-starter/internal/router"
)

// We don't actually use API key, but OpenAPI v2 enforces this way
// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization

func main() {

	if err := config.LoadConfig(".env"); err != nil {
		log.Fatalf("Loading application config failed: %v", err)
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	if err := config.ConnectRedis(); err != nil {
		log.Fatalf("Cache DB connection failed: %v", err)
	}
	defer config.Config.RedisClient.Close()

	r := router.RegisterRoutes()
	router.SetupSwagger(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Config.ServerPort),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}
	log.Println("Graceful shutdown finished.")
}
