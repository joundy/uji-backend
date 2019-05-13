package exam

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//Repository represent course repository contract
type Repository interface {
	FetchG(*models.Filter) ([]*models.ExamG, error)
	GetByID(*primitive.ObjectID) (*models.Exam, error)
	UpdateByID(*primitive.ObjectID, *models.Exam) error
	DeleteByID(*primitive.ObjectID) error
	Create(*models.Exam) (*models.ResID, error)
}
