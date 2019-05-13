package question

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/question"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type questionUsecase struct {
	qRepository question.Repository
}

//NewQuestionUsecase represent initializatin NewQuestionUsecase
func NewQuestionUsecase(eR question.Repository) Usecase {
	return &questionUsecase{eR}
}

func (qU *questionUsecase) GetByID(ID *primitive.ObjectID) (*models.Question, error) {
	question, err := qU.qRepository.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (qU *questionUsecase) UpdateByID(ID *primitive.ObjectID, question *models.Question) error {
	err := qU.qRepository.UpdateByID(ID, question)
	if err != nil {
		return err
	}

	return nil
}

func (qU *questionUsecase) DeleteByID(ID *primitive.ObjectID) error {
	err := qU.qRepository.DeleteByID(ID)
	if err != nil {
		return err
	}

	return nil
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
