package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type generateHandler struct {
	// eGUsecase examgroup.Usecase
}

//NewExamLogsHandler represent initialization generateHandler
func NewExamLogHandler(e *echo.Echo) {
	handler := &generateHandler{}

	g := e.Group("/examLogs/generate")

	g.POST("", handler.Generate)
}

func (eGH *generateHandler) Generate(eC echo.Context) error {
	return eC.JSON(http.StatusOK, "OK")
}
