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
		log.Fatalln(`Please make sure .env file exists or env variable TWELVE_FACTOR_MODE is set to "true": `, err)
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalln("Database connection failed: ", err)
	}

	if err := config.ConnectRedis(); err != nil {
		log.Fatalln("Cache DB connection failed: ", err)
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
		log.Fatal("HTTP server shutdown failed: ", err)
	}
	log.Println("Graceful shutdown finished.")
}
