package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//JWTClaims represent model for jwtClaims data
type JWTClaims struct {
	ID         primitive.ObjectID
	Fullname   string
	Email      string
	UserTypeID primitive.ObjectID
	jwt.StandardClaims
}

//ResToken represent model for ResToken data
type ResToken struct {
	Token string `json:"token"`
}
