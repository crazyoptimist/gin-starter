package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gin-starter/internal/core/config"
	"gin-starter/internal/core/logger"
	"gin-starter/internal/core/router"
)

// We are not using an API key here, but in OpenAPI v2 there is no better way to configure this.
// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization

func main() {

	if err := config.LoadConfig(".env"); err != nil {
		log.Println("Warning: dotenv file is missing, please make sure you have configured environment variables properly", err)
	}

	config.ConnectDB()

	logger.InitAppLogger()

	r := router.RegisterRoutes()
	router.SetupSwagger(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Config.ServerPort),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("ListenAndServe: ", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
