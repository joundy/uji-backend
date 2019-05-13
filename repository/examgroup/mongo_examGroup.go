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

func (m *mongoExamGroupRepository) GetByID(ID *primitive.ObjectID) (*models.ExamGroup, error) {
	collection := m.mgoClient.Database("uji").Collection("examGroups")

	var examGroup models.ExamGroup

	err := collection.FindOne(context.TODO(), bson.D{{"_id", ID}}).Decode(&examGroup)

	if err != nil {
		return nil, err
	}

	return &examGroup, nil
}

func (m *mongoExamGroupRepository) UpdateByID(ID *primitive.ObjectID, examGroup *models.ExamGroup) error {
	collection := m.mgoClient.Database("uji").Collection("examGroups")

	_, err := collection.UpdateOne(context.TODO(), bson.D{
		{"_id", ID},
	}, bson.D{
		{"$set", examGroup},
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoExamGroupRepository) DeleteByID(ID *primitive.ObjectID) error {
	collection := m.mgoClient.Database("uji").Collection("examGroups")

	_, err := collection.DeleteOne(
		context.Background(),
		bson.D{
			{"_id", ID},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoExamGroupRepository) Create(eG *models.ExamGroup) (*models.ResID, error) {
	collection := m.mgoClient.Database("uji").Collection("examGroups")

	_, err := collection.InsertOne(context.Background(), eG)
	if err != nil {
		return nil, err
	}

	resID := models.ResID{
		ID: eG.ID,
	}

	return &resID, nil
}

func (m *mongoExamGroupRepository) FetchG(mF *models.Filter) ([]*models.ExamGroupG, error) {
	collection := m.mgoClient.Database("uji").Collection("examGroups")
	var examGroupGs []*models.ExamGroupG

	fBLevel := bson.D{{"$match", bson.D{}}}
	if mF.Level != "" {
		fBLevel = bson.D{{"$match", bson.D{{"level", mF.Level}}}}
	}

	fBClass := bson.D{{"$match", bson.D{}}}
	if mF.Class != "" {
		fBClass = bson.D{{"$match", bson.D{{"class", mF.Class}}}}
	}

	fBTag := bson.D{{"$match", bson.D{}}}
	if mF.Tag != "" {
		fBTag = bson.D{{"$match", bson.D{{"tag", mF.Tag}}}}
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$sort", bson.D{
				{"_id", -1},
			}},
		},
		fBLevel,
		fBClass,
		fBTag,
		// bson.D{
		// 	{"$lookup", bson.D{
		// 		{"from", "levels"},
		// 		{"localField", "levelId"},
		// 		{"foreignField", "_id"},
		// 		{"as", "level"},
		// 	}},
		// },
		// bson.D{
		// 	{"$lookup", bson.D{
		// 		{"from", "classes"},
		// 		{"localField", "classId"},
		// 		{"foreignField", "_id"},
		// 		{"as", "class"},
		// 	}},
		// },
		// bson.D{
		// 	{"$addFields", bson.D{
		// 		{"level", bson.D{
		// 			{"$arrayElemAt", []interface{}{"$level", 0}},
		// 		}},
		// 		{"class", bson.D{
		// 			{"$arrayElemAt", []interface{}{"$class", 0}},
		// 		}},
		// 	}},
		// },
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
