package user

import (
	"context"

	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoUserRepository struct {
	mgoClient *mongo.Client
}

//NewMongoUserRespository  represent initialitation mongoUserRepository
func NewMongoUserRespository(c *mongo.Client) Repository {
	return &mongoUserRepository{c}
}

func (m *mongoUserRepository) GetByEmail(email string) (*models.User, error) {
	collection := m.mgoClient.Database("uji").Collection("users")

	var user models.User

	err := collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
