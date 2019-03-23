package examlog

import (
	"fmt"

	"github.com/haffjjj/uji-backend/repository/question"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/exam"
	"github.com/haffjjj/uji-backend/repository/examlog"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type examLogUsecase struct {
	eLRepository examlog.Repository
	eRepository  exam.Repository
	qRepository  question.Repository
}

//NewExamLogUsecase represent initializatin NewExamLogUsecase
func NewExamLogUsecase(eLR examlog.Repository, eR exam.Repository, qR question.Repository) Usecase {
	return &examLogUsecase{eLR, eR, qR}
}

func (c *examLogUsecase) GetByID(i primitive.ObjectID) (*models.ExamLog, error) {
	examLog, err := c.eLRepository.GetByID(i)

	if err != nil {
		return nil, err
	}

	return examLog, nil
}

func (c *examLogUsecase) Generate(userID, examID primitive.ObjectID) error {

	// exam, err := c.eRepository.GetByID(examID)
	// if err != nil {
	// 	return err
	// }

	filter := models.Filter{Start: 0, Limit: 100}
	questionGs, err := c.qRepository.FetchG(filter)
	if err != nil {
		return err
	}

	questionG := *questionGs[0]

	fmt.Println(questionG.Data)

	return nil
}
