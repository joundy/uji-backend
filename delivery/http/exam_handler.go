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
	c.POST("", handler.Create)
	c.GET("/:id", handler.GetByID)
	c.PUT("/:id", handler.UpdateByID)
	c.DELETE("/:id", handler.DeleteByID)
}

func (eH *examHandler) GetByID(eC echo.Context) error {
	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	exam, err := eH.eUsecase.GetByID(&ID)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, exam)
}

func (eH *examHandler) UpdateByID(eC echo.Context) error {
	var exam models.Exam
	eC.Bind(&exam)

	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	err = eH.eUsecase.UpdateByID(&ID, &exam)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (eH *examHandler) DeleteByID(eC echo.Context) error {
	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	err = eH.eUsecase.DeleteByID(&ID)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (eH *examHandler) Create(eC echo.Context) error {
	var exam models.Exam
	eC.Bind(&exam)

	exam.ID = primitive.NewObjectID()

	if examGroupIDP, ok := eC.QueryParams()["examGroupId"]; ok {
		examGroupIDHex, err := primitive.ObjectIDFromHex(examGroupIDP[0])
		if err != nil {
			return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		}
		exam.ExamGroupID = examGroupIDHex
	}

	resID, err := eH.eUsecase.Create(&exam)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, resID)
}

func (eH *examHandler) FetchG(eC echo.Context) error {
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

	if examGroupIDP, ok := eC.QueryParams()["examGroupId"]; ok {
		examGroupIDHex, err := primitive.ObjectIDFromHex(examGroupIDP[0])
		if err != nil {
			return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		}
		mF.ExamGroupID = examGroupIDHex
	}

	if examGroupSlugP, ok := eC.QueryParams()["examGroupSlug"]; ok {
		mF.ExamGroupSlug = examGroupSlugP[0]
	}

	//usecase
	examsGs, err := eH.eUsecase.FetchG(&mF)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examsGs)
}
