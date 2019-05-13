package question

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Usecase represent course usecase contract
type Usecase interface {
	Create(*models.Question) (*models.ResID, error)
	GetByID(*primitive.ObjectID) (*models.Question, error)
	UpdateByID(*primitive.ObjectID, *models.Question) error
	DeleteByID(*primitive.ObjectID) error
	FetchG(*models.Filter) ([]*models.QuestionG, error)
}
