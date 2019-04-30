package examgroup

import "github.com/haffjjj/uji-backend/models"

//Usecase represent examgroup usecase contract
type Usecase interface {
	FetchG(*models.Filter) ([]*models.ExamGroupG, error)
}
