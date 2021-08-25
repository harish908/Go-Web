package handlers

import (
	"PortalClient/configs"
	"PortalClient/internal/data"
	"PortalClient/internal/tracing"
	"PortalClient/internal/utils"
	"context"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func GetIdeasHandler(w http.ResponseWriter, r *http.Request) {
	// Start a span for request
	span := tracing.StartSpanFromRequest(configs.GetTracer(), r)
	// set tags and logs based on
	//span.SetTag("URL :", r.URL)
	//span.LogFields(
	//	log.string("event", "getIdeas"))
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	body, err := data.GetIdeas(ctx)

	if err != nil {
		log.Error(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = utils.ToJSON(w, body)
	if err != nil {
		log.Error("Unable to marshal json", err)
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func PostIdeaHandler(w http.ResponseWriter, r *http.Request) {
	// Start a span for request
	span := tracing.StartSpanFromRequest(configs.GetTracer(), r)
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read data", http.StatusInternalServerError)
	}

	response, err := data.PostIdea(body, ctx)
	if err != nil {
		log.Error(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = utils.ToJSON(w, response)
	if err != nil {
		log.Error("Unable to marshal json", err)
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
