package course

import "github.com/haffjjj/uji-backend/models"

//Usecase represent course usecase contract
type Usecase interface {
	FetchG(*models.Filter) ([]*models.CourseG, error)
}
