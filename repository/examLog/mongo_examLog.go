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

func (m *mongoExamLogRepository) Start(IDHex, userIDHex *primitive.ObjectID, e *models.ExamLog) error {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	_, err := collection.UpdateOne(context.TODO(), bson.D{
		{"_id", IDHex},
		{"userId", userIDHex},
	}, bson.D{
		{"$set", bson.D{
			{"isStart", e.IsStart},
			{"startTime", e.StartTime},
			{"endTime", e.EndTime},
		}},
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoExamLogRepository) Submit(IDHex, userIDHex *primitive.ObjectID, e *models.ExamLog) error {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	_, err := collection.UpdateOne(context.TODO(), bson.D{
		{"_id", IDHex},
		{"userId", userIDHex},
	}, bson.D{
		{"$set", bson.D{
			{"result", e.Result},
			{"isSubmit", e.IsSubmit},
			{"timeSpent", e.TimeSpent},
		}},
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoExamLogRepository) SetAnswer(IDHex, userIDHex, questionIDHex *primitive.ObjectID, selectedIdsHex *[]primitive.ObjectID) error {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	_, err := collection.UpdateOne(context.TODO(), bson.D{
		{"_id", IDHex},
		{"userId", userIDHex},
		{"questions._id", questionIDHex},
	}, bson.D{
		{"$set", bson.D{
			{"questions.$.answer.selectedIds", selectedIdsHex},
		}},
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoExamLogRepository) SetIsMarked(IDHex, userIDHex, questionIDHex *primitive.ObjectID, isMarked *bool) error {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	_, err := collection.UpdateOne(context.TODO(), bson.D{
		{"_id", IDHex},
		{"userId", userIDHex},
		{"questions._id", questionIDHex},
	}, bson.D{
		{"$set", bson.D{
			{"questions.$.isMarked", isMarked},
		}},
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoExamLogRepository) GetByID(IDHex, userIDHex *primitive.ObjectID) (*models.ExamLog, error) {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	var examLog models.ExamLog

	err := collection.FindOne(context.TODO(), bson.D{{"_id", IDHex}, {"userId", userIDHex}}).Decode(&examLog)

	if err != nil {
		return nil, err
	}

	return &examLog, nil
}

func (m *mongoExamLogRepository) Store(e *models.ExamLog) (*models.ExamLog, error) {
	collection := m.mgoClient.Database("uji").Collection("examLogs")

	res, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		return nil, err
	}
	e.ID = res.InsertedID.(primitive.ObjectID)

	return e, nil
}

func (m *mongoExamLogRepository) FetchG(userIDHex *primitive.ObjectID, mF *models.Filter) ([]*models.ExamLogG, error) {
	collection := m.mgoClient.Database("uji").Collection("examLogs")
	var examLogGs []*models.ExamLogG

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"userId", userIDHex},
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
