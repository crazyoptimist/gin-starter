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
	"gin-starter/internal/logger"
	"gin-starter/internal/router"
)

// We don't use API key, but in OpenAPI v2 there's no better way to configure this.
// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization

func main() {

	if err := config.LoadConfig(".env"); err != nil {
		log.Fatalln(`Please make sure dotenv file exists or env variable TWELVE_FACTOR_MODE is set to "true": `, err)
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalln("Error: database connection failed: ", err)
	}

	appLogger, err := logger.InitAppLogger()
	if err != nil {
		log.Fatalln("Error: logger initialization failed: ", err)
	}
	defer appLogger.Instance.Sync()

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
	quit := make(chan os.Signal, 1)
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
