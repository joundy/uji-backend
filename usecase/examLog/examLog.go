package examlog

import (
	"math/rand"
	"time"

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

	exam, err := c.eRepository.GetByID(examID)
	if err != nil {
		return err
	}

	filter := models.Filter{Start: 0, Limit: 100}
	questionGs, err := c.qRepository.FetchG(filter)
	if err != nil {
		return err
	}

	questionG := questionGs[0]
	qDataRaw := questionG.Data

	if exam.IsRandom == true {
		shuffleQuestions(&qDataRaw)
	}

	if exam.MaxQuestion > len(qDataRaw) {
		exam.MaxQuestion = len(qDataRaw)
	}

	examLog := models.ExamLog{
		UserID: userID,
		ExamID: examID,
		Exam: models.ExamLogExam{
			Title:        exam.Title,
			Description:  exam.Description,
			Duration:     exam.Duration,
			Source:       exam.Source,
			Point:        exam.Point,
			PassingGrade: exam.PassingGrade,
		},
		IsSubmit:  false,
		Questions: qDataRaw[:exam.MaxQuestion],
	}

	err = c.eLRepository.Store(&examLog)
	if err != nil {
		return err
	}

	return nil
}

func (c *examLogUsecase) FetchG(f models.Filter) ([]*models.ExamLogG, error) {
	examLogGs, err := c.eLRepository.FetchG(f)

	if err != nil {
		return nil, err
	}

	return examLogGs, nil
}

func shuffleQuestions(q *[]models.Question) {
	rand.Seed(time.Now().UnixNano())

	for i := len(*q) - 1; i >= 0; i-- {
		r := rand.Intn(len(*q) - 1)

		(*q)[i], (*q)[r] = (*q)[r], (*q)[i]
	}
}
