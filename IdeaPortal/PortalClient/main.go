package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harish908/Go-Web/IdeaPortal/PortalClient/config"
	"log"
	"net/http"
	"time"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Initialize config file
	config.ConfigFile()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", test)
	srv := http.Server{
		Handler: router,
		Addr:    "127.0.0.1:80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
