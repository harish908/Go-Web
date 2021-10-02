package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request){
	err := json.NewEncoder(w).Encode("Portal-Server running")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}