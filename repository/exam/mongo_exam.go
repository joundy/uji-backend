package exam

import (
	"context"
	"errors"

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

func (m *mongoExamRepository) Create(e *models.Exam) (*models.ResID, error) {
	collection := m.mgoClient.Database("uji").Collection("exams")

	_, err := collection.InsertOne(context.Background(), e)
	if err != nil {
		return nil, err
	}

	resID := models.ResID{
		ID: e.ID,
	}

	return &resID, nil
}

func (m *mongoExamRepository) GetByID(i primitive.ObjectID) (*models.Exam, error) {
	collection := m.mgoClient.Database("uji").Collection("exams")

	var exams []*models.Exam

	// var exam models.Exam

	// err := collection.FindOne(context.TODO(), bson.D{{"_id", i}}).Decode(&exam)

	// if err != nil {
	// 	return nil, err
	// }

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{{"$match", bson.D{{"_id", i}}}},
		bson.D{
			{"$lookup", bson.D{
				{"from", "examGroups"},
				{"localField", "examGroupId"},
				{"foreignField", "_id"},
				{"as", "examGroup"},
			}},
		},
		bson.D{
			{"$addFields", bson.D{
				{"examGroup", bson.D{
					{"$arrayElemAt", []interface{}{"$examGroup", 0}},
				}},
			}},
		},
	})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var elem models.Exam
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		exams = append(exams, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	if len(exams) == 0 || exams == nil {
		return nil, errors.New("mongo: no documents in result")
	}

	return exams[0], nil
}

func (m *mongoExamRepository) FetchG(mF *models.Filter) ([]*models.ExamG, error) {
	collection := m.mgoClient.Database("uji").Collection("exams")
	var examGs []*models.ExamG

	var zHex primitive.ObjectID

	fBExamGroupId := bson.D{{"$match", bson.D{}}}
	if mF.ExamGroupID != zHex {
		fBExamGroupId = bson.D{{"$match", bson.D{{"examGroupId", mF.ExamGroupID}}}}
	}

	fBExamGroupSlug := bson.D{{"$match", bson.D{}}}
	if mF.ExamGroupSlug != "" {
		fBExamGroupSlug = bson.D{{"$match", bson.D{{"examGroup.slug", mF.ExamGroupSlug}}}}
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$sort", bson.D{
				{"_id", -1},
			}},
		},
		fBExamGroupId,
		bson.D{
			{"$lookup", bson.D{
				{"from", "examGroups"},
				{"localField", "examGroupId"},
				{"foreignField", "_id"},
				{"as", "examGroup"},
			}},
		},
		bson.D{
			{"$addFields", bson.D{
				{"examGroup", bson.D{
					{"$arrayElemAt", []interface{}{"$examGroup", 0}},
				}},
			}},
		},
		fBExamGroupSlug,
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
