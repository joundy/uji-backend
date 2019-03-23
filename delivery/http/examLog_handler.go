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
