package examgroup

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Usecase represent examgroup usecase contract
type Usecase interface {
	FetchG(*models.Filter) ([]*models.ExamGroupG, error)
	GetByID(*primitive.ObjectID) (*models.ExamGroup, error)
	UpdateByID(*primitive.ObjectID, *models.ExamGroup) error
	DeleteByID(*primitive.ObjectID) error
	Create(*models.ExamGroup) (*models.ResID, error)
}
