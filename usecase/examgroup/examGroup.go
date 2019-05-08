package examgroup

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/examgroup"
)

type examGroupUsecase struct {
	cRepository examgroup.Repository
}

//NewExamGroupUsecase represent initializatin courseUsecase
func NewExamGroupUsecase(cR examgroup.Repository) Usecase {
	return &examGroupUsecase{cR}
}

func (c *examGroupUsecase) Create(eG *models.ExamGroup) (*models.ResID, error) {
	resID, err := c.cRepository.Create(eG)
	if err != nil {
		return nil, err
	}

	return resID, nil
}

func (c *examGroupUsecase) FetchG(mF *models.Filter) ([]*models.ExamGroupG, error) {
	examGroupGs, err := c.cRepository.FetchG(mF)

	if err != nil {
		return nil, err
	}

	return examGroupGs, nil
}
