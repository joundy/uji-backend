package http

import (
	"net/http"
	"strconv"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/exam"
	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//examGroupHandler represent handler for course
type examHandler struct {
	eUsecase exam.Usecase
}

//NewExamHandler represent initialization courseHandler
func NewExamHandler(e *echo.Echo, eU exam.Usecase) {
	handler := &examHandler{eU}

	c := e.Group("/exams")

	c.GET("", handler.FetchG)
}

func (eH *examHandler) FetchG(eC echo.Context) error {
	filter := models.Filter{Start: 0, Limit: 100}

	if startP, ok := eC.QueryParams()["start"]; ok {
		start, err := strconv.Atoi(startP[0])
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		filter.Start = start
	}

	if limitP, ok := eC.QueryParams()["limit"]; ok {
		limit, err := strconv.Atoi(limitP[0])
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		filter.Limit = limit
	}

	if examGroupIDP, ok := eC.QueryParams()["examGroup"]; ok {
		examGroupIDHex, err := primitive.ObjectIDFromHex(examGroupIDP[0])
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		filter.ExamGroupID = examGroupIDHex
	}

	examsGs, err := eH.eUsecase.FetchG(filter)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examsGs)
}
