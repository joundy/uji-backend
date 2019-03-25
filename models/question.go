package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Question is represent model for Question data
type Question struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Answer Answer             `json:"anwers" bson:"answer"`
}

//Answer is represent model for Answer data
type Answer struct {
	List        []AnswerList         `json:"list" bson:"list"`
	CorrectIds  []primitive.ObjectID `json:"-" bson:"correctIds"`
	SelectedIds []primitive.ObjectID `json:"selectedIds" form:"selectedId" bson:"selectedIds"`
}

//AnswerList is represent model for AnswerList data
type AnswerList struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" bson:"title"`
}

//QuestionG is represent model for QuestionG data
type QuestionG struct {
	Data  []Question `json:"data" bson:"data"`
	Count int        `json:"count" bson:"count"`
}
