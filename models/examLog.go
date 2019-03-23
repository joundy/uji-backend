package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//ExamLog is represent model for ExamLog data
type ExamLog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	StartTime time.Time          `json:"startTime" bson:"startTime"`
	EndTime   time.Time          `json:"endTime" bson:"endTime"`
	Result    ExamLogResult      `json:"result" bson:"result"`
	IsSubmit  bool               `json:"isSubmit" bson:"isSubmit"`
	Questions []Question         `json:"questions" bson:"questions"`
}

//ExamLogResult is represent model for ExamLogResult data
type ExamLogResult struct {
	Pass   int `json:"pass" bson:"pass"`
	Failed int `json:"failed" bson:"failed"`
}

//ExamLogG is represent model for ExamLogG data
type ExamLogG struct {
	Data  []Class `json:"data" bson:"data"`
	Count int     `json:"count" bson:"count"`
}
