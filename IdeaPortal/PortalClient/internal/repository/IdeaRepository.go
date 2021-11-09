package repository

import (
	"PortalClient/internal/models"
	"PortalClient/configs"
	"context"

	"github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
	log "github.com/sirupsen/logrus"
)

func GetIdeas(ctx context.Context) ([]models.Idea, error) {
	db := configs.GetMySqlDB()
	span, _ := opentracing.StartSpanFromContext(ctx, "getIdeas")
	defer span.Finish()

	var ideas []models.Idea
	query := "SELECT * FROM Idea"
	span.SetTag("db.statement", query)

	result, err := db.Query(query)
	if err != nil {
		span.SetTag("error", true)
		span.LogFields(otlog.Error(err))
		log.Error("Unable to fetch data ", err)
		return nil, err
	}
	for result.Next() {
		var idea models.Idea
		err = result.Scan(&idea.Id, &idea.Title, &idea.Description, &idea.EstimatedTime, &idea.CreatedDate)
		if err != nil {
			log.Error("Unable to scan data ", err)
			return nil, err
		}
		ideas = append(ideas, idea)
	}
	return ideas, nil
}

func PostIdea(idea *models.Idea) bool {
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
