package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
	"v1/configs"
	"v1/internal/middleware"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Print("test")
	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Initialize config file
	configs.InitConfigFile()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.loggingMiddleware)
	router.HandleFunc("/test", test)
	srv := &http.Server{
		Handler: router,
		Addr:    viper.GetString("ApiInfo.PortalClient"),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(srv.ListenAndServe())
}
