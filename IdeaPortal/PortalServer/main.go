package main

import (
	"PortalServer/configs"
	"PortalServer/internal/middleware"
	"PortalServer/internal/routes"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func init() {
	// Initialize configs
	configs.InitConfigFile()
	configs.InitTracer()
	configs.InitDB()

	environment := viper.GetString("Environment")
	if environment != "Development" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
	// router handler
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggingMiddleware)
	routes.HandleRoutes(router)

	// Create and start server
	log.Info("PortalServer starting")
	srv := &http.Server{
		Handler: router,
		Addr:    ":8001",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(srv.ListenAndServe())
}
