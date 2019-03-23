package http

import (
	"net/http"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/examlog"
	"github.com/labstack/echo"
)

type examLogHandler struct {
	eGUsecase examlog.Usecase
}

//NewExamLogHandler represent initialization generateHandler
func NewExamLogHandler(e *echo.Echo, eLU examlog.Usecase) {
	handler := &examLogHandler{eLU}

	g := e.Group("/examLogs")

	g.POST("", handler.Generate)
	g.GET("/:id", handler.GetByID)
}

func (eGH *examLogHandler) Generate(eC echo.Context) error {
	return eC.JSON(http.StatusOK, "OK")
}

func (eGH *examLogHandler) GetByID(eC echo.Context) error {
	examLog, err := eGH.eGUsecase.GetByID("sdf")

	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examLog)
}
