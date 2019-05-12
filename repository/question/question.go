package question

import (
	"context"

	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoQuestionRepository struct {
	mgoClient *mongo.Client
}

//NewMongoQuestionRepository represent initialization NewMongoQuestionRepository
func NewMongoQuestionRepository(c *mongo.Client) Repository {
	return &mongoQuestionRepository{c}
}

func (m *mongoQuestionRepository) Create(q *models.Question) (*models.ResID, error) {
	collection := m.mgoClient.Database("uji").Collection("questions")

	_, err := collection.InsertOne(context.Background(), q)
	if err != nil {
		return nil, err
	}

	resID := models.ResID{
		ID: q.ID,
	}

	return &resID, nil
}

func (m *mongoQuestionRepository) FetchG(mF *models.Filter) ([]*models.QuestionG, error) {
	collection := m.mgoClient.Database("uji").Collection("questions")
	var questionGs []*models.QuestionG
	var zHex primitive.ObjectID

	fBExamId := bson.D{{"$match", bson.D{}}}
	if mF.ExamID != zHex {
		fBExamId = bson.D{{"$match", bson.D{{"examId", mF.ExamID}}}}
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$sort", bson.D{
				{"_id", -1},
			}},
		},
		fBExamId,
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

		var elem models.QuestionG
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		questionGs = append(questionGs, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	if questionGs == nil {
		return []*models.QuestionG{}, nil
	}

	return questionGs, nil
}
