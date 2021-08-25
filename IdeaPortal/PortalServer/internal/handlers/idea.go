package handlers

import (
	"PortalServer/configs"
	"PortalServer/internal/data"
	"PortalServer/internal/tracing"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetIdeasHandler(w http.ResponseWriter, r *http.Request) {
	span := tracing.StartSpanFromRequest(configs.GetTracer(), r)
	defer span.Finish()
	ideas := data.GetIdeas()

	e := json.NewEncoder(w)
	err := e.Encode(ideas)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func PostIdeaHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read data", http.StatusInternalServerError)
	}

	response := data.PostIdea(body)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
