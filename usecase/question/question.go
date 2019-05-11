package question

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/question"
)

type questionUsecase struct {
	qRepository question.Repository
}

//NewQuestionUsecase represent initializatin NewQuestionUsecase
func NewQuestionUsecase(eR question.Repository) Usecase {
	return &questionUsecase{eR}
}

func (qU *questionUsecase) Create(q *models.Question) (*models.ResID, error) {
	resID, err := qU.qRepository.Create(q)
	if err != nil {
		return nil, err
	}

	return resID, nil
}

func (qU *questionUsecase) FetchG(f *models.Filter) ([]*models.QuestionG, error) {
	questionGs, err := qU.qRepository.FetchG(f)
	if err != nil {
		return nil, err
	}

	return questionGs, nil
}
