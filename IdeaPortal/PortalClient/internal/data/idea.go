package data

import (
	"PortalClient/internal/utils"
	"context"
)

type Idea struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	EstimatedTime int `json:"estimatedTime"`
	CreatedDate string `json:"createdData"`
}

type Ideas []*Idea

func GetIdeas(ctx context.Context) ([]byte, error){
	apiData := make(chan []byte)
	apiErr := make(chan error)
	go utils.Ping("ideas", "GET", nil, apiData, apiErr, ctx, "getIdeas")
	//log.Info("Line executes before ping function, since we used go routines")
	return <-apiData, <-apiErr
}

func PostIdea(body []byte, ctx context.Context) ([]byte, error){
	apiData := make(chan []byte)
	apiErr := make(chan error)
	go utils.Ping("postIdea", "POST", body, apiData, apiErr, ctx, "postIdea")
	//log.Info("Line executes before ping function, since we used go routines")
	return <-apiData, <-apiErr
}