package exam

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/exam"
)

type examUsecase struct {
	eRepository exam.Repository
}

//NewExamUsecase represent initializatin courseUsecase
func NewExamUsecase(eR exam.Repository) Usecase {
	return &examUsecase{eR}
}

func (eU *examUsecase) Create(e *models.Exam) (*models.ResID, error) {
	resID, err := eU.eRepository.Create(e)
	if err != nil {
		return nil, err
	}

	return resID, nil
}

func (eU *examUsecase) FetchG(mF *models.Filter) ([]*models.ExamG, error) {
	examsGs, err := eU.eRepository.FetchG(mF)

	if err != nil {
		return nil, err
	}

	return examsGs, nil
}
