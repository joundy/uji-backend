package exam

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/exam"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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

func (eU *examUsecase) GetByID(ID *primitive.ObjectID) (*models.Exam, error) {
	exam, err := eU.eRepository.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return exam, nil
}

func (eU *examUsecase) UpdateByID(ID *primitive.ObjectID, exam *models.Exam) error {
	err := eU.eRepository.UpdateByID(ID, exam)
	if err != nil {
		return err
	}

	return nil
}

func (eU *examUsecase) DeleteByID(ID *primitive.ObjectID) error {
	err := eU.eRepository.DeleteByID(ID)
	if err != nil {
		return err
	}

	return nil
}

func (eU *examUsecase) FetchG(mF *models.Filter) ([]*models.ExamG, error) {
	examsGs, err := eU.eRepository.FetchG(mF)

	if err != nil {
		return nil, err
	}

	return examsGs, nil
}
