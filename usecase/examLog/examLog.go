package examlog

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/examlog"
)

type examLogUsecase struct {
	eLRepository examlog.Repository
}

//NewExamLogUsecase represent initializatin NewExamLogUsecase
func NewExamLogUsecase(cR examlog.Repository) Usecase {
	return &examLogUsecase{cR}
}

func (c *examLogUsecase) GetByID(i string) (*models.ExamLog, error) {
	examLog, err := c.eLRepository.GetByID(i)

	if err != nil {
		return nil, err
	}

	return examLog, nil
}
