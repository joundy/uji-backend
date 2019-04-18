package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Filter represent model for Filter data
type Filter struct {
	Start         int
	Limit         int
	Tag           string
	ExamGroupSlug string
	Class         string
	Level         string
	ExamGroupID   primitive.ObjectID
	ExamID        primitive.ObjectID
}
