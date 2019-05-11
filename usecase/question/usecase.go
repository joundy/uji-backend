package question

import (
	"github.com/haffjjj/uji-backend/models"
)

//Usecase represent course usecase contract
type Usecase interface {
	Create(*models.Question) (*models.ResID, error)
	FetchG(*models.Filter) ([]*models.QuestionG, error)
}
