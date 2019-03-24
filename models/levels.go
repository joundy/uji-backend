package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Level is represent model for course data
type Level struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
}

//LevelG is represent model for courseGroup data
type LevelG struct {
	Data  []Level `json:"data" bson:"data"`
	Count int     `json:"count" bson:"count"`
}
