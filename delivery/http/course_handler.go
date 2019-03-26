package http

import (
	"net/http"
	"strconv"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/course"
	"github.com/labstack/echo"
)

type courseHandler struct {
	cUsecase course.Usecase
}

//NewCourseHandler represent initialization courseHandler
func NewCourseHandler(e *echo.Echo, cU course.Usecase) {
	handler := &courseHandler{cU}

	c := e.Group("/courses")

	c.GET("", handler.FetchG)
}

//FetchG is method from courseHandler
func (cH *courseHandler) FetchG(eC echo.Context) error {

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

	//usecase
	courseGs, err := cH.cUsecase.FetchG(&mF)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, courseGs)
}
