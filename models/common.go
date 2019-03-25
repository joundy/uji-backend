package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

//ResID is represent model for ResID data
type ResID struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
