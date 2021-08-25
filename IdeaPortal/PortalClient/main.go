package main

import (
	"PortalClient/configs"
	"PortalClient/internal/middleware"
	"PortalClient/internal/routes"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func init() {
	// Initialize config file
	configs.InitConfigFile()

	environment := viper.GetString("Environment")
	if environment != "Development" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// Initialize tracer
	configs.InitTracer()
}

func main() {
	// route handler
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggingMiddleware)
	routes.HandleRoutes(router)

	// Create and start server
	log.Info("PortalClient starting")
	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(srv.ListenAndServe())
}
