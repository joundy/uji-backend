package examlog

import (
	"context"
	"fmt"

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

func (m *mongoExamLogRepository) Submit(IDHex, userIDHex *primitive.ObjectID) error {
	fmt.Println(IDHex, userIDHex)

	return nil
}

func (m *mongoExamLogRepository) SetAnswer(IDHex, userIDHex, questionIDHex *primitive.ObjectID, isSelectedIdsHex *[]primitive.ObjectID) error {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	_, err := collection.UpdateOne(context.TODO(), bson.D{
		{"_id", IDHex},
		{"userId", userIDHex},
		{"questions._id", questionIDHex},
	}, bson.D{
		{"$set", bson.D{
			{"questions.$.answer.selectedIds", isSelectedIdsHex},
		}},
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoExamLogRepository) GetByID(i *primitive.ObjectID) (*models.ExamLog, error) {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	var examLog models.ExamLog

	err := collection.FindOne(context.TODO(), bson.D{{"_id", i}}).Decode(&examLog)

	if err != nil {
		return nil, err
	}

	return &examLog, nil
}

func (m *mongoExamLogRepository) Store(e *models.ExamLog) error {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	_, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoExamLogRepository) FetchG(mF models.Filter) ([]*models.ExamLogG, error) {
	collection := m.mgoClient.Database("uji").Collection("examLogs")
	var examLogGs []*models.ExamLogG

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"userId", mF.UserID},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"questions", 0},
			}},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"count", bson.D{{"$sum", 1}}},
				{"data", bson.D{{"$push", "$$ROOT"}}},
			}},
		},
		bson.D{
			{"$unwind", "$data"},
		},
		bson.D{
			{"$replaceRoot", bson.D{
				{"newRoot", bson.D{
					{"$mergeObjects", bson.A{"$data", "$$ROOT"}},
				}},
			}},
		},
		bson.D{
			{"$skip", mF.Start},
		},
		bson.D{
			{"$limit", mF.Limit},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"count", bson.D{
					{"$first", "$count"},
				}},
				{"data", bson.D{
					{"$push", "$data"},
				}},
			}},
		},
	})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem models.ExamLogG
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		examLogGs = append(examLogGs, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	if examLogGs == nil {
		return []*models.ExamLogG{}, nil
	}

	return examLogGs, nil
}
