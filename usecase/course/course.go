package course

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/course"
)

type courseUsecase struct {
	cRepository course.Repository
}

//NewCourseUsecase represent initializatin courseUsecase
func NewCourseUsecase(cR course.Repository) Usecase {
	return &courseUsecase{cR}
}

func (c *courseUsecase) FetchG(mF *models.Filter) ([]*models.CourseG, error) {
	courseGs, err := c.cRepository.FetchG(mF)

	if err != nil {
		return nil, err
	}

	return courseGs, nil
}
