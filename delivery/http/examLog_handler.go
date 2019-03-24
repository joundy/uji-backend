package http

import (
	"net/http"

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
	g.GET("/:id", handler.GetByID)
	g.GET("", handler.FetchG)
}

func (eGH *examLogHandler) FetchG(eC echo.Context) error {
	mF := models.Filter{Start: 0, Limit: 100}

	// if startP, ok := eC.QueryParams()["start"]; ok {
	// 	start, err := strconv.Atoi(startP[0])
	// 	if err != nil {
	// 		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	// 	}
	// 	filter.Start = start
	// }

	// if limitP, ok := eC.QueryParams()["limit"]; ok {
	// 	limit, err := strconv.Atoi(limitP[0])
	// 	if err != nil {
	// 		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	// 	}
	// 	filter.Limit = limit
	// }

	// if examGroupIDP, ok := eC.QueryParams()["examGroup"]; ok {
	// 	examGroupIDHex, err := primitive.ObjectIDFromHex(examGroupIDP[0])
	// 	if err != nil {
	// 		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	// 	}
	// 	filter.ExamGroupID = examGroupIDHex
	// }

	examLogGs, err := eGH.eGUsecase.FetchG(mF)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examLogGs)
}

func (eGH *examLogHandler) Generate(eC echo.Context) error {

	examIDF := eC.FormValue("examId")
	examIDHex, err := primitive.ObjectIDFromHex(examIDF)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex("5c94d2b450e8986339d26534")
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	err = eGH.eGUsecase.Generate(userIDHex, examIDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, "OK")
}

func (eGH *examLogHandler) GetByID(eC echo.Context) error {

	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	examLog, err := eGH.eGUsecase.GetByID(IDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examLog)
}
