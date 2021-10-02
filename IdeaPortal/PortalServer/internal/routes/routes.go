package routes

import (
	"PortalServer/internal/handlers"
	"github.com/gorilla/mux"
)

func HandleRoutes(router *mux.Router){
	router.HandleFunc("/healthCheck", handlers.HealthCheckHandler)
	router.HandleFunc("/ideas", handlers.GetIdeasHandler)
	router.HandleFunc("/postIdea", handlers.PostIdeaHandler).Methods("POST")
}
