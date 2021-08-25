package data

import (
	"PortalServer/configs"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// Declare struct based on rules from ultimate-go course
type Idea struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	EstimatedTime int    `json:"estimatedTime"`
	CreatedDate   string `json:"createdData"`
}

func GetIdeas() []Idea {
	db := configs.GetMySqlDB()
	var ideas []Idea
	result, err := db.Query("SELECT * FROM Idea")
	if err != nil {
		log.Error("Unable to fetch data ", err)
		return nil
	}
	for result.Next() {
		var idea Idea
		err = result.Scan(&idea.Id, &idea.Title, &idea.Description, &idea.EstimatedTime, &idea.CreatedDate)
		if err != nil {
			log.Error("Unable to scan data ", err)
			return nil
		}
		ideas = append(ideas, idea)
	}
	return ideas
}

func PostIdea(body []byte) bool {
	db := configs.GetMySqlDB()
	var idea Idea
	err := json.Unmarshal(body, &idea)
	if err != nil {
		log.Error("Unable to marshal json ", err)
		return false
	}

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
