package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//ExamLog is represent model for ExamLog data
type ExamLog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	ExamID    primitive.ObjectID `json:"examId" bson:"examId"`
	Exam      ExamLogExam        `json:"exam" bson:"exam"`
	Questions []Question         `json:"questions" bson:"questions"`
	StartTime time.Time          `json:"startTime" bson:"startTime"`
	EndTime   time.Time          `json:"endTime" bson:"endTime"`
	Result    ExamLogResult      `json:"result" bson:"result"`
	IsSubmit  bool               `json:"isSubmit" bson:"isSubmit"`
}

//ExamLogResult is represent model for ExamLogResult data
type ExamLogResult struct {
	Pass   int `json:"pass" bson:"pass"`
	Failed int `json:"failed" bson:"failed"`
}

//ExamLogExam is represent model for ExamLogExam data
type ExamLogExam struct {
	Title        string `json:"title" bson:"title"`
	Description  string `json:"description" bson:"description"`
	Duration     int    `json:"duration" bson:"duration"`
	Source       string `json:"source" bson:"source"`
	Point        int    `json:"point" bson:"point"`
	PassingGrade int    `json:"passingGrade" bson:"passingGrade"`
}

//ExamLogG is represent model for ExamLogG data
type ExamLogG struct {
	Data  []Class `json:"data" bson:"data"`
	Count int     `json:"count" bson:"count"`
}