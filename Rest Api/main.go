package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(body, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(Articles)
}

func singleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["title"]
	for _, article := range Articles {
		if article.Title == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/articles", allArticles)
	router.HandleFunc("/singleArticle/{title}", singleArticle)
	router.HandleFunc("/article", createArticle).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
