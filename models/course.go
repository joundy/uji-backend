package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Course is represent model for course data
type Course struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	// Deadline    time.Time `json:"deadline" bson:"deadline"`
}

//CourseG is represent model for courseGroup data
type CourseG struct {
	Data  []Course `json:"data" bson:"data"`
	Count int      `json:"count" bson:"count"`
}
