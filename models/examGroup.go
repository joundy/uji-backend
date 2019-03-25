package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//ExamGroup is represent model for course data
type ExamGroup struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CourseID    primitive.ObjectID `json:"courseId" bson:"courseId"`
	LevelID     primitive.ObjectID `json:"-" bson:"levelId"`
	ClassID     primitive.ObjectID `json:"-" bson:"classId"`
	Description string             `json:"description" bson:"description"`
	Level       Level              `json:"level" bson:"level"`
	Class       Class              `json:"class" bson:"class"`
}

//ExamGroupG is represent model for ExamGroupGroup data
type ExamGroupG struct {
	Data  []ExamGroup `json:"data" bson:"data"`
	Count int         `json:"count" bson:"count"`
}
