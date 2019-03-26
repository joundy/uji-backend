package course

import (
	"context"

	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoCourseRepository struct {
	mgoClient *mongo.Client
}

//NewMongoCourseRepository represent initialization mongoCourseRepository
func NewMongoCourseRepository(c *mongo.Client) Repository {
	return &mongoCourseRepository{c}
}

func (m *mongoCourseRepository) FetchG(mF *models.Filter) ([]*models.CourseG, error) {
	collection := m.mgoClient.Database("uji").Collection("courses")
	var courseG []*models.CourseG

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
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

		var elem models.CourseG
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		courseG = append(courseG, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	if courseG == nil {
		return []*models.CourseG{}, nil
	}

	return courseG, nil
}
