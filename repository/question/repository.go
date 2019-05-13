package question

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Repository represent course repository contract
type Repository interface {
	FetchG(*models.Filter) ([]*models.QuestionG, error)
	GetByID(*primitive.ObjectID) (*models.Question, error)
	UpdateByID(*primitive.ObjectID, *models.Question) error
	DeleteByID(*primitive.ObjectID) error
	Create(*models.Question) (*models.ResID, error)
}
