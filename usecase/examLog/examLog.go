package examlog

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/examlog"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type examLogUsecase struct {
	eLRepository examlog.Repository
}

//NewExamLogUsecase represent initializatin NewExamLogUsecase
func NewExamLogUsecase(cR examlog.Repository) Usecase {
	return &examLogUsecase{cR}
}

func (c *examLogUsecase) GetByID(i primitive.ObjectID) (*models.ExamLog, error) {
	examLog, err := c.eLRepository.GetByID(i)

	if err != nil {
		return nil, err
	}

	return examLog, nil
}

func (c *examLogUsecase) Generate(userID, examID primitive.ObjectID) error {

	// fmt.Println(userID, examID)

	return nil
}
