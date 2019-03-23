package exam

import (
	"context"

	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoExamRepository struct {
	mgoClient *mongo.Client
}

//NewMongoExamRepository represent initialization mongoCourseRepository
func NewMongoExamRepository(c *mongo.Client) Repository {
	return &mongoExamRepository{c}
}

func (m *mongoExamRepository) FetchG(mF models.Filter) ([]*models.ExamG, error) {
	collection := m.mgoClient.Database("uji").Collection("exams")
	var examGs []*models.ExamG

	var zHex primitive.ObjectID

	fBExamGroupId := bson.D{{"$match", bson.D{}}}
	if mF.ExamGroupID != zHex {
		fBExamGroupId = bson.D{{"$match", bson.D{{"examGroupId", mF.ExamGroupID}}}}
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		fBExamGroupId,
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

		var elem models.ExamG
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		examGs = append(examGs, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	if examGs == nil {
		return []*models.ExamG{}, nil
	}

	return examGs, nil
}
