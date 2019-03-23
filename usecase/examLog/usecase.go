package examlog

import "github.com/haffjjj/uji-backend/models"

//Usecase represent course usecase contract
type Usecase interface {
	GetByID(i string) (*models.ExamLog, error)
}
