package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//User is represent model for User data
type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Fullname   string             `json:"fullname" bson:"fullname"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	Schools    string             `json:"schools" bson:"schools"`
	Point      int                `json:"point" bson:"point"`
	UserTypeID primitive.ObjectID `json:"userTypeId" bson:"userTypeId"`
}
