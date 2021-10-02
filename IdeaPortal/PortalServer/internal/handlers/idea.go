package handlers

import (
	"PortalServer/configs"
	"PortalServer/internal/tracing"
	"context"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	ottag "github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// Declare struct based on rules from ultimate-go course
type Idea struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	EstimatedTime int `json:"estimatedTime"`
	CreatedDate string `json:"createdData"`
}

func GetIdeasHandler(w http.ResponseWriter, r *http.Request){
	// Start a span for request
	span := tracing.StartSpanFromRequest(r)
	// set tags
	// ottag.SpanKindRPCClient.Set(span)
	ottag.HTTPUrl.Set(span, r.URL.Path)
	ottag.HTTPMethod.Set(span, r.Method)

	// Inorder to capture endTime of request, we must call finish()
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(r.Context(), span)
	ideas, err := GetIdeas(ctx)
	if err!=nil{
		span.SetTag("error", true)
		span.LogFields(otlog.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	err = e.Encode(ideas)

	// We can write response model using this format
	span.LogKV(
		"test1", "testing",
		"test2", "testing",
	)
	if err!=nil{
		span.SetTag("error", true)
		span.LogFields(otlog.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span.SetTag("response", "Successfully fetched ideas from database")
}

func PostIdeaHandler(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var idea *Idea
	err = json.Unmarshal(body, &idea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := idea.PostIdea()
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}




func GetIdeas(ctx context.Context) ([]Idea, error){
	db := configs.GetMySqlDB()
	span, _ := opentracing.StartSpanFromContext(ctx, "getIdeas")
	defer span.Finish()

	var ideas []Idea
	query := "SELECT * FROM Idea"
	span.SetTag("db.statement", query)

	result, err := db.Query(query)
	if err != nil {
		span.SetTag("error", true)
		span.LogFields(otlog.Error(err))
		log.Error("Unable to fetch data ", err)
		return nil, err
	}
	for result.Next(){
		var idea Idea
		err = result.Scan(&idea.Id, &idea.Title, &idea.Description, &idea.EstimatedTime, &idea.CreatedDate)
		if err != nil{
			log.Error("Unable to scan data ", err)
			return nil, err
		}
		ideas = append(ideas, idea)
	}
	return ideas, nil
}

func (idea *Idea) PostIdea() bool {
	db := configs.GetMySqlDB()
	ins, err := db.Prepare("INSERT INTO Idea (Title, Description, EstimatedTime, CreatedDate) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Error("Error in prepare statement ", err)
		return false
	}
	defer ins.Close()
	_, err = ins.Exec(idea.Title, idea.Description, idea.EstimatedTime, idea.CreatedDate)
	if err != nil {
		log.Error("Error while inserting data ", err)
		return false
	}
	return true
}
