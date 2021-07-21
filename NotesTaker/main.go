package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Note struct {
	Text string `json:"Text"`
}

func addNote(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var note Note
	json.Unmarshal(body, &note)
	log.Print("Data request: ", note)
}

// Run "fresh" cmd in terminal for golang hot reload
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addNote", addNote).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
