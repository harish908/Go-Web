package routes

import (
	"PortalClient/internal/handlers"
	"github.com/gorilla/mux"
)

func HandleRoutes(router *mux.Router) {
	router.HandleFunc("/ideas", handlers.GetIdeasHandler)
	router.HandleFunc("/postIdea", handlers.PostIdeaHandler).Methods("POST")
}
