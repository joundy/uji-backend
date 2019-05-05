package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Exam is represent model for course data
type Exam struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ExamGroupID  primitive.ObjectID `json:"examGroupId" bson:"examGroupId"`
	ExamGroup    ExamGroup          `json:"-" bson:"examGroup"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Slug         string             `json:"slug" bson:"slug"`
	Duration     int                `json:"duration" bson:"duration"`
	Source       string             `json:"source" bson:"source"`
	IsRandom     bool               `json:"-" bson:"isRandom"`
	MaxQuestion  int                `json:"maxQuestion" bson:"maxQuestion"`
	Point        int                `json:"point" bson:"point"`
	PassingGrade int                `json:"passingGrade" bson:"passingGrade"`
	Questions    []Question         `json:"-" bson:"questions"`
}

//ExamG is represent model for courseGroup data
type ExamG struct {
	Data  []Exam `json:"data" bson:"data"`
	Count int    `json:"count" bson:"count"`
}
