package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Class is represent model for course data
type Class struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
}

//ClassG is represent model for courseGroup data
type ClassG struct {
	Data  []Class `json:"data" bson:"data"`
	Count int     `json:"count" bson:"count"`
}
