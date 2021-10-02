package main

import (
	"PortalServer/configs"
	"PortalServer/internal/middleware"
	"PortalServer/internal/routes"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func init(){
	// Initialize configs
	configs.InitConfigFile()
	configs.InitTracer()
	configs.InitDB()

	environment := viper.GetString("Environment")
	if environment != "Development"{
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
	// router handler
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggingMiddleware)
	routes.HandleRoutes(router)

	// Profiler
	// Run program and run "pprof -web http://localhost:8001/debug/pprof/heap"
	// router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	// Create and start server
	log.Info("PortalServer starting")
	srv := &http.Server{
		Handler: router,
		Addr: ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(srv.ListenAndServe())

	// Close tracer
	configs.CloseTracer()


}
