package examlog

import (
	"context"

	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoExamLogRepository struct {
	mgoClient *mongo.Client
}

//NewMongoExamLogRepository represent initialization mongoCourseRepository
func NewMongoExamLogRepository(c *mongo.Client) Repository {
	return &mongoExamLogRepository{c}
}

func (m *mongoExamLogRepository) GetByID(i primitive.ObjectID) (*models.ExamLog, error) {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	var examLog models.ExamLog

	err := collection.FindOne(context.TODO(), bson.D{{"_id", i}}).Decode(&examLog)

	if err != nil {
		return nil, err
	}

	return &examLog, nil
}
