package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Question is represent model for Question data
type Question struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ExamID   primitive.ObjectID `json:"examId" bson:"examId,omitempty"`
	Title    string             `json:"title" bson:"title"`
	Answer   Answer             `json:"answer" bson:"answer"`
	IsMarked bool               `json:"isMarked" bson:"isMarked"`
}

//Answer is represent model for Answer data
type Answer struct {
	List        []AnswerList         `json:"list" bson:"list"`
	SelectedIds []primitive.ObjectID `json:"selectedIds" form:"selectedId" bson:"selectedIds"`
}

//AnswerList is represent model for AnswerList data
type AnswerList struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	IsCorrect bool               `json:"isCorrect" bson:"isCorrect"`
}

//QuestionG is represent model for QuestionG data
type QuestionG struct {
	Data  []Question `json:"data" bson:"data"`
	Count int        `json:"count" bson:"count"`
}
