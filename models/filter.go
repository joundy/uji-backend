package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Filter represent model for Filter data
type Filter struct {
	Start       int
	Limit       int
	ClassID     primitive.ObjectID
	LevelID     primitive.ObjectID
	ExamGroupID primitive.ObjectID
	ExamID      primitive.ObjectID
	CourseID    primitive.ObjectID
}
