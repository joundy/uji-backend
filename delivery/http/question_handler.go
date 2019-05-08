package http

import (
	"net/http"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/question"
	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//examGroupHandler represent handler for course
type questionHandler struct {
	qUsecase question.Usecase
}

//NewQuestionHandler represent initialization NewQuestionHandler
func NewQuestionHandler(e *echo.Echo, qU question.Usecase) {
	handler := &questionHandler{qU}

	c := e.Group("/questions")

	c.POST("", handler.Create)
}

func (qH *questionHandler) Create(eC echo.Context) error {
	var question models.Question
	eC.Bind(&question)

	question.ID = primitive.NewObjectID()

	if examIDP, ok := eC.QueryParams()["examId"]; ok {
		examIDHex, err := primitive.ObjectIDFromHex(examIDP[0])
		if err != nil {
			return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		}
		question.ExamID = examIDHex
	}

	for i := range question.Answer.List {
		question.Answer.List[i].ID = primitive.NewObjectID()
	}

	question.Answer.SelectedIds = []primitive.ObjectID{}

	resID, err := qH.qUsecase.Create(&question)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, resID)
}
