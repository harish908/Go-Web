package handlers

import (
	"PortalClient/internal/models"
	"PortalClient/internal/repository"
	"PortalClient/pkg/tracing"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/opentracing/opentracing-go"
	ottag "github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
)

// @Summary Get all Ideas
// @Tags Ideas
// @Accept json
// @Produce json
// @Success 200 {object} models.Idea
// @Router /api/ideas [get]
func GetIdeasHandler(w http.ResponseWriter, r *http.Request) {
	// Start a span for request
	span := tracing.StartSpanFromRequest(r)
	// set tags
	// ottag.SpanKindRPCClient.Set(span)
	ottag.HTTPUrl.Set(span, r.URL.Path)
	ottag.HTTPMethod.Set(span, r.Method)

	// Inorder to capture endTime of request, we must call finish()
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(r.Context(), span)
	// try to use goroutines here
	ideas, err := repository.GetIdeas(ctx)
	if err != nil {
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
	if err != nil {
		span.SetTag("error", true)
		span.LogFields(otlog.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span.SetTag("response", "Successfully fetched ideas from database")
}

// @Summary Post Idea
// @Tags Ideas
// @Accept json
// @Produce json
// @Param request body models.Idea true "Post ideas"
// @Success 200 {string} string "success"
// @Router /api/postIdea [post]
func PostIdeaHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var idea *models.Idea
	err = json.Unmarshal(body, &idea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := repository.PostIdea(idea)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
// @Summary Get all Ideas
// @Tags Ideas
// @Accept json
// @Produce json
// @Success 200 {object} data.Idea
// @Router /api/ideas [get]
func GetIdeasHandler(w http.ResponseWriter, r *http.Request) {
	// Start a span for request
	span := tracing.StartSpanFromRequest(r)
	tracing.LogURL(span, r.URL.Path, r.Method)

	// Inorder to capture endTime of request, we must call finish()
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(r.Context(), span)

	body, err := gateways.GetIdeas(ctx)
	if err != nil {
		tracing.LogError(span, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// We can write response model using this format
	span.LogKV(
		"test1", "testing",
		"test2", "testing",
	)
	span.SetTag("response", "Successfully fetched ideas")
	_, err = w.Write(body)
	if err != nil {
		tracing.LogError(span, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Post Idea
// @Tags Ideas
// @Accept json
// @Produce json
// @Param request body data.Idea true "Post ideas"
// @Success 200 {string} string "success"
// @Router /api/postIdea [post]
func PostIdeaHandler(w http.ResponseWriter, r *http.Request) {
	// Start a span for request
	span := tracing.StartSpanFromRequest(r)
	tracing.LogURL(span, r.URL.Path, r.Method)

	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read data", http.StatusInternalServerError)
	}

	response, err := gateways.PostIdea(body, ctx)
	if err != nil {
		tracing.LogError(span, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = serializer.ToJSON(w, response)
	if err != nil {
		tracing.LogError(span, err)
		log.Error("Unable to marshal json", err)
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
*/
