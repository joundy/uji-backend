package examlog

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Usecase represent course usecase contract
type Usecase interface {
	GetByID(i *primitive.ObjectID) (*models.ExamLog, error)
	Generate(userID, examID primitive.ObjectID) error
	FetchG(models.Filter) ([]*models.ExamLogG, error)
	SetAnswer(IDHex, userIDHex, questionIDHex *primitive.ObjectID, isSelectedIdsHex *[]primitive.ObjectID) error
	Submit(IDHex, userIDHex *primitive.ObjectID) error
}
