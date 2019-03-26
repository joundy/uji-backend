package http

import (
	"net/http"
	"strconv"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/examlog"
	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type examLogHandler struct {
	eGUsecase examlog.Usecase
}

//NewExamLogHandler represent initialization generateHandler
func NewExamLogHandler(e *echo.Echo, eLU examlog.Usecase) {
	handler := &examLogHandler{eLU}

	g := e.Group("/examLogs")

	g.POST("/generate", handler.Generate)
	g.GET("/:id", handler.GetByIDAndStart)
	g.GET("", handler.FetchG)
	g.PUT("/:id/setAnswers/questions/:questionId", handler.SetAnswer)
	g.POST("/:id/submit", handler.Submit)
}

func (eLH *examLogHandler) FetchG(eC echo.Context) error {
	mF := models.Filter{Start: 0, Limit: 100}

	if startP, ok := eC.QueryParams()["start"]; ok {
		start, err := strconv.Atoi(startP[0])
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		mF.Start = start
	}

	if limitP, ok := eC.QueryParams()["limit"]; ok {
		limit, err := strconv.Atoi(limitP[0])
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		mF.Limit = limit
	}

	userIDHex, err := primitive.ObjectIDFromHex("5c94d2b450e8986339d26534")
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	examLogGs, err := eLH.eGUsecase.FetchG(&userIDHex, &mF)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examLogGs)
}

func (eLH *examLogHandler) Generate(eC echo.Context) error {

	examIDF := eC.FormValue("examId")
	examIDHex, err := primitive.ObjectIDFromHex(examIDF)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex("5c94d2b450e8986339d26534")
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	resID, err := eLH.eGUsecase.Generate(userIDHex, examIDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, resID)
}

func (eLH *examLogHandler) GetByIDAndStart(eC echo.Context) error {

	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex("5c94d2b450e8986339d26534")
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	examLog, err := eLH.eGUsecase.GetByIDAndStart(&IDHex, &userIDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examLog)
}

func (eLH *examLogHandler) SetAnswer(eC echo.Context) error {
	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	questionIDP := eC.Param("questionId")
	questionIDHex, err := primitive.ObjectIDFromHex(questionIDP)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex("5c94d2b450e8986339d26534")
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	var isSelectedIdsHex []primitive.ObjectID

	fParams, _ := eC.FormParams()
	isSelectedIds := fParams["isSelectedId"]

	for _, v := range isSelectedIds {
		elemHex, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}

		isSelectedIdsHex = append(isSelectedIdsHex, elemHex)
	}

	err = eLH.eGUsecase.SetAnswer(&IDHex, &userIDHex, &questionIDHex, &isSelectedIdsHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (eLH *examLogHandler) Submit(eC echo.Context) error {

	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex("5c94d2b450e8986339d26534")
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	err = eLH.eGUsecase.Submit(&IDHex, &userIDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}
