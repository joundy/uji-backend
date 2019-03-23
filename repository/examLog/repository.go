package examlog

import (
	"github.com/haffjjj/uji-backend/models"
)

//Repository represent course repository contract
type Repository interface {
	GetByID(i string) (*models.ExamLog, error)
}
