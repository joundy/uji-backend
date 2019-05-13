package http

import (
	"net/http"
	"strconv"

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
	c.GET("", handler.FetchG)
	c.GET("/:id", handler.GetByID)
	c.PUT("/:id", handler.UpdateByID)
	c.DELETE("/:id", handler.DeleteByID)
}

func (qH *questionHandler) GetByID(eC echo.Context) error {
	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	question, err := qH.qUsecase.GetByID(&ID)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, question)
}

func (qH *questionHandler) UpdateByID(eC echo.Context) error {
	var question models.Question
	eC.Bind(&question)

	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	for i := range question.Answer.List {
		question.Answer.List[i].ID = primitive.NewObjectID()
	}

	question.Answer.SelectedIds = []primitive.ObjectID{}

	err = qH.qUsecase.UpdateByID(&ID, &question)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (qH *questionHandler) DeleteByID(eC echo.Context) error {
	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	err = qH.qUsecase.DeleteByID(&ID)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
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

func (qH *questionHandler) FetchG(eC echo.Context) error {
	mF := models.Filter{Start: 0, Limit: 100}

	if startP, ok := eC.QueryParams()["start"]; ok {
		start, err := strconv.Atoi(startP[0])
		if err != nil {
			return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		}
		mF.Start = start
	}

	if limitP, ok := eC.QueryParams()["limit"]; ok {
		limit, err := strconv.Atoi(limitP[0])
		if err != nil {
			return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		}
		mF.Limit = limit
	}

	if examIDP, ok := eC.QueryParams()["examId"]; ok {
		examIDHex, err := primitive.ObjectIDFromHex(examIDP[0])
		if err != nil {
			return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		}
		mF.ExamID = examIDHex
	}

	//usecase
	questionsGs, err := qH.qUsecase.FetchG(&mF)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, questionsGs)
}
