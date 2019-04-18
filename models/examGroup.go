package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//ExamGroup is represent model for course data
type ExamGroup struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Level       string             `json:"level" bson:"level"`
	Class       string             `json:"class" bson:"class"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Slug        string             `json:"slug" bson:"slug"`
	Tag         string             `json:"tag" bson:"tag"`
}

//ExamGroupG is represent model for ExamGroupGroup data
type ExamGroupG struct {
	Data  []ExamGroup `json:"data" bson:"data"`
	Count int         `json:"count" bson:"count"`
}
