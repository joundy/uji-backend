package examlog

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Repository represent course repository contract
type Repository interface {
	GetByID(IDHex, userIDHex *primitive.ObjectID) (*models.ExamLog, error)
	Store(e *models.ExamLog) error
	FetchG(userIDHex *primitive.ObjectID, mF models.Filter) ([]*models.ExamLogG, error)
	SetAnswer(IDHex, userIDHex, questionIDHex *primitive.ObjectID, isSelectedIdsHex *[]primitive.ObjectID) error
	Submit(IDHex, userIDHex *primitive.ObjectID, e *models.ExamLog) error
}
