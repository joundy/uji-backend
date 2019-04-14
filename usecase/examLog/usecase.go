package examlog

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Usecase represent course usecase contract
type Usecase interface {
	GetByIDAndStart(IDHex, userIDHex *primitive.ObjectID) (*models.ExamLog, error)
	Generate(userID, examID primitive.ObjectID) (*models.ResID, error)
	FetchG(userIDHex *primitive.ObjectID, mF *models.Filter) ([]*models.ExamLogG, error)
	SetAnswer(IDHex, userIDHex, questionIDHex *primitive.ObjectID, selectedIdsHex *[]primitive.ObjectID) error
	SetIsMarked(IDHex, userIDHex, questionIDHex *primitive.ObjectID, isMarked *bool) error
	Submit(IDHex, userIDHex *primitive.ObjectID) error
}
