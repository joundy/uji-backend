package question

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/question"
)

type questionUsecase struct {
	qRepository question.Repository
}

//NewExamUsecase represent initializatin courseUsecase
func NewQuestionUsecase(eR question.Repository) Usecase {
	return &questionUsecase{eR}
}

func (eU *questionUsecase) Create(e *models.Question) (*models.ResID, error) {
	// resID, err := eU.eRepository.Create(e)
	// if err != nil {
	// 	return nil, err
	// }

	// return resID, nil

	return nil, nil
}
