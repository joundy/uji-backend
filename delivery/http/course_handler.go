package http

import (
	"net/http"

	"github.com/haffjjj/uji-backend/usecase/course"
	"github.com/labstack/echo"
)

//CourseHandler represent handler for course
type CourseHandler struct {
	cUsecase course.Usecase
}

//NewTagHandler represent initialization courseHandler
func NewTagHandler(e *echo.Echo, cU course.Usecase) {
	handler := &CourseHandler{cU}

	c := e.Group("/courses")

	c.GET("", handler.FetchG)
}

//FetchG is method from courseHandler
func (cH *CourseHandler) FetchG(eC echo.Context) error {
	return eC.JSON(http.StatusOK, "OK")
}
