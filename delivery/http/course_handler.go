package http

import (
	"net/http"
	"strconv"

	"github.com/haffjjj/uji-backend/models"
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

	courseGs, err := cH.cUsecase.FetchG(filter)

	if err != nil {
		eC.JSON(http.StatusInternalServerError, "Error")
	}

	return eC.JSON(http.StatusOK, courseGs)
}
