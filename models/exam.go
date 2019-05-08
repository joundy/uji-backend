package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Exam is represent model for exam data
type Exam struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ExamGroupID  primitive.ObjectID `json:"examGroupId" bson:"examGroupId,omitempty"`
	ExamGroup    ExamGroup          `json:"-" bson:"examGroup,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Slug         string             `json:"slug" bson:"slug"`
	Duration     int                `json:"duration" bson:"duration"`
	Source       string             `json:"source" bson:"source"`
	IsRandom     bool               `json:"-" bson:"isRandom"`
	MaxQuestion  int                `json:"maxQuestion" bson:"maxQuestion"`
	Point        int                `json:"point" bson:"point"`
	PassingGrade int                `json:"passingGrade" bson:"passingGrade"`
}

//ExamInput is represent model for examInput data
type ExamInput struct {
	ExamGroupID  string `json:"examGroupId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Slug         string `json:"slug"`
	Duration     int    `json:"duration"`
	Source       string `json:"source"`
	IsRandom     bool   `json:"-"`
	MaxQuestion  int    `json:"maxQuestion"`
	Point        int    `json:"point"`
	PassingGrade int    `json:"passingGrade"`
}

//ExamG is represent model for courseGroup data
type ExamG struct {
	Data  []Exam `json:"data" bson:"data"`
	Count int    `json:"count" bson:"count"`
}
