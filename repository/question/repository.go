package question

import (
	"github.com/haffjjj/uji-backend/models"
)

//Repository represent course repository contract
type Repository interface {
	FetchG(models.Filter) ([]*models.QuestionG, error)
}
