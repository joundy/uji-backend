package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Question is represent model for Question data
type Question struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Title   string             `json:"title" bson:"title"`
	Answers []ListAnswer       `json:"anwers" bson:"answers"`
}

//ListAnswer is represent model for List data
type ListAnswer struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	IsCorrect  bool               `json:"isCorrect" bson:"isCorrect"`
	IsSelected bool               `json:"isSelected" bson:"isSelected"`
}

//QuestionG is represent model for QuestionG data
type QuestionG struct {
	Data  []Question `json:"data" bson:"data"`
	Count int        `json:"count" bson:"count"`
}
