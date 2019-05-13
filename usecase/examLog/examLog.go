package examlog

import (
	"errors"
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

func (c *examLogUsecase) Submit(IDHex, userIDHex *primitive.ObjectID) error {
	examLog, err := c.eLRepository.GetByID(IDHex, userIDHex)
	if err != nil {
		return err
	}

	if examLog.IsSubmit == true {
		return errors.New("exam already submitted")
	}

	if examLog.IsStart == false {
		return errors.New("the exam hasn't started yet, please start first")
	}

	//includes function
	contains := func(selectedIds []primitive.ObjectID, id primitive.ObjectID) bool {
		for _, v := range selectedIds {
			if v == id {
				return true
			}
		}
		return false
	}

	//check list answer one by one
	eqAnswers := func(list []models.AnswerList, selectedIds []primitive.ObjectID) bool {
		for _, v := range list {
			if v.IsCorrect == true {
				if contains(selectedIds, v.ID) != true {
					return false
				}
			}
		}
		return true
	}

	for _, v := range examLog.Questions {
		if eqAnswers(v.Answer.List, v.Answer.SelectedIds) == true {
			examLog.Result.Pass++
		} else {
			examLog.Result.Failed++
		}
	}

	examLog.TimeSpent = time.Now().Local().Sub(examLog.StartTime).Seconds()

	examLog.IsSubmit = true

	err = c.eLRepository.Submit(IDHex, userIDHex, examLog)
	if err != nil {
		return err
	}

	return nil
}

func (c *examLogUsecase) SetAnswer(IDHex, userIDHex, questionIDHex *primitive.ObjectID, selectedIdsHex *[]primitive.ObjectID) error {
	examLog, err := c.eLRepository.GetByID(IDHex, userIDHex)
	if err != nil {
		return err
	}

	if examLog.IsSubmit == true {
		return errors.New("exam already submitted")
	}

	if examLog.IsStart == false {
		return errors.New("the exam hasn't started yet, please start first")
	}

	if time.Now().Local().Sub(examLog.EndTime).Seconds() >= 0 {
		return errors.New("timeout, please submit exam..")
	}

	err = c.eLRepository.SetAnswer(IDHex, userIDHex, questionIDHex, selectedIdsHex)
	if err != nil {
		return err
	}

	return nil
}

func (c *examLogUsecase) SetIsMarked(IDHex, userIDHex, questionIDHex *primitive.ObjectID, isMarked *bool) error {
	examLog, err := c.eLRepository.GetByID(IDHex, userIDHex)
	if err != nil {
		return err
	}

	if examLog.IsSubmit == true {
		return errors.New("exam already submitted")
	}

	if examLog.IsStart == false {
		return errors.New("the exam hasn't started yet, please start first")
	}

	if time.Now().Local().Sub(examLog.EndTime).Seconds() >= 0 {
		return errors.New("timeout, please submit exam..")
	}

	err = c.eLRepository.SetIsMarked(IDHex, userIDHex, questionIDHex, isMarked)
	if err != nil {
		return err
	}

	return nil
}

func (c *examLogUsecase) GetByIDAndStart(IDHex, userIDHex *primitive.ObjectID) (*models.ExamLog, error) {
	examLog, err := c.eLRepository.GetByID(IDHex, userIDHex)
	if err != nil {
		return nil, err
	}

	if examLog.IsStart == false {
		examLog.IsStart = true
		examLog.StartTime = time.Now()
		examLog.EndTime = time.Now().Local().Add(time.Second * time.Duration(examLog.Exam.Duration))

		err = c.eLRepository.Start(IDHex, userIDHex, examLog)
		if err != nil {
			return nil, err
		}
	}

	if examLog.IsSubmit == false {
		for i, v := range examLog.Questions {
			for j := range v.Answer.List {
				examLog.Questions[i].Answer.List[j].IsCorrect = false
			}
		}
	}

	return examLog, nil
}

func (c *examLogUsecase) Generate(userIDHex, examIDHex primitive.ObjectID) (*models.ResID, error) {

	exam, err := c.eRepository.GetByID(&examIDHex)
	if err != nil {
		return nil, err
	}

	mF := models.Filter{Start: 0, Limit: 200, ExamID: examIDHex}
	questionGs, err := c.qRepository.FetchG(&mF)
	if err != nil {
		return nil, err
	}

	if len(questionGs) == 0 {
		return nil, errors.New("No question found")
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
		UserID: userIDHex,
		ExamID: exam.ID,
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

	res, err := c.eLRepository.Store(&examLog)
	if err != nil {
		return nil, err
	}

	resID := models.ResID{ID: res.ID}

	return &resID, nil
}

func (c *examLogUsecase) FetchG(userIDHex *primitive.ObjectID, mF *models.Filter) ([]*models.ExamLogG, error) {
	examLogGs, err := c.eLRepository.FetchG(userIDHex, mF)

	if err != nil {
		return nil, err
	}

	return examLogGs, nil
}

func shuffleQuestions(q *[]models.Question) {
	if len(*q) > 1 {
		rand.Seed(time.Now().UnixNano())
		for i := len(*q) - 1; i >= 0; i-- {
			r := rand.Intn(len(*q) - 1)

			(*q)[i], (*q)[r] = (*q)[r], (*q)[i]
		}
	}
}
