package examgroup

import (
	"context"

	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoExamGroupRepository struct {
	mgoClient *mongo.Client
}

//NewMongoExamGroupRepository represent initialization mongoCourseRepository
func NewMongoExamGroupRepository(c *mongo.Client) Repository {
	return &mongoExamGroupRepository{c}
}

func (m *mongoExamGroupRepository) FetchG(mF models.Filter) ([]*models.ExamGroupG, error) {
	collection := m.mgoClient.Database("uji").Collection("examGroups")
	var examGroupGs []*models.ExamGroupG

	filterByCourseId := bson.D{{"$match", bson.D{}}}

	if mF.CourseID != "" {
		CourseIDHex, err := primitive.ObjectIDFromHex(mF.CourseID)
		if err != nil {
			return nil, err
		}
		filterByCourseId = bson.D{{"$match", bson.D{{"courseId", CourseIDHex}}}}
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		filterByCourseId,
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

		var elem models.ExamGroupG
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		examGroupGs = append(examGroupGs, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	if examGroupGs == nil {
		return []*models.ExamGroupG{}, nil
	}

	return examGroupGs, nil
}
