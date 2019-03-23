package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Exam is represent model for course data
type Exam struct {
	CourseID     primitive.ObjectID `json:"courseId" bson:"courseId"`
	ExamGroupID  primitive.ObjectID `json:"examGroupId" bson:"examGroupId"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Duration     int                `json:"duration" bson:"duration"`
	Source       string             `json:"source" bson:"source"`
	IsRandom     bool               `json:"isRandom" bson:"deadline"`
	MaxQuestion  int                `json:"maxQuestion" bson:"maxQuestion"`
	Point        int                `json:"point" bson:"point"`
	PassingGrade int                `json:"passingGrade" bson:"passingGrade"`
}

//ExamG is represent model for courseGroup data
type ExamG struct {
	Data  []Exam `json:"data" bson:"data"`
	Count int    `json:"count" bson:"count"`
}
