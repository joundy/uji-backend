package examgroup

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/examgroup"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type examGroupUsecase struct {
	cRepository examgroup.Repository
}

//NewExamGroupUsecase represent initializatin courseUsecase
func NewExamGroupUsecase(cR examgroup.Repository) Usecase {
	return &examGroupUsecase{cR}
}

func (c *examGroupUsecase) GetByID(ID *primitive.ObjectID) (*models.ExamGroup, error) {
	examGroup, err := c.cRepository.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return examGroup, nil
}

func (c *examGroupUsecase) UpdateByID(ID *primitive.ObjectID, examGroup *models.ExamGroup) error {
	err := c.cRepository.UpdateByID(ID, examGroup)
	if err != nil {
		return err
	}

	return nil
}

func (c *examGroupUsecase) DeleteByID(ID *primitive.ObjectID) error {
	err := c.cRepository.DeleteByID(ID)
	if err != nil {
		return err
	}

	return nil
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
