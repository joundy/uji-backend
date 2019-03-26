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

func (m *mongoExamGroupRepository) FetchG(mF *models.Filter) ([]*models.ExamGroupG, error) {
	collection := m.mgoClient.Database("uji").Collection("examGroups")
	var examGroupGs []*models.ExamGroupG

	var zHex primitive.ObjectID

	fBCourseId := bson.D{{"$match", bson.D{}}}
	if mF.CourseID != zHex {
		fBCourseId = bson.D{{"$match", bson.D{{"courseId", mF.CourseID}}}}
	}

	fBLevelId := bson.D{{"$match", bson.D{}}}
	if mF.LevelID != zHex {
		fBLevelId = bson.D{{"$match", bson.D{{"levelId", mF.LevelID}}}}
	}

	fBClassId := bson.D{{"$match", bson.D{}}}
	if mF.ClassID != zHex {
		fBClassId = bson.D{{"$match", bson.D{{"classId", mF.ClassID}}}}
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		fBCourseId,
		fBLevelId,
		fBClassId,
		// bson.D{
		// 	{"$lookup", bson.D{
		// 		{"from", "courses"},
		// 		{"localField", "courseId"},
		// 		{"foreignField", "_id"},
		// 		{"as", "course"},
		// 	}},
		// },
		bson.D{
			{"$lookup", bson.D{
				{"from", "levels"},
				{"localField", "levelId"},
				{"foreignField", "_id"},
				{"as", "level"},
			}},
		},
		bson.D{
			{"$lookup", bson.D{
				{"from", "classes"},
				{"localField", "classId"},
				{"foreignField", "_id"},
				{"as", "class"},
			}},
		},
		bson.D{
			{"$addFields", bson.D{
				// {"course", bson.D{
				// 	{"$arrayElemAt", []interface{}{"$course", 0}},
				// }},
				{"level", bson.D{
					{"$arrayElemAt", []interface{}{"$level", 0}},
				}},
				{"class", bson.D{
					{"$arrayElemAt", []interface{}{"$class", 0}},
				}},
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
