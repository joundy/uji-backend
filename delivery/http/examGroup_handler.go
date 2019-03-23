package http

import (
	"net/http"
	"strconv"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/examgroup"
	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//examGroupHandler represent handler for course
type examGroupHandler struct {
	eGUsecase examgroup.Usecase
}

//NewExamGroupHandler represent initialization courseHandler
func NewExamGroupHandler(e *echo.Echo, eGU examgroup.Usecase) {
	handler := &examGroupHandler{eGU}

	c := e.Group("/examGroups")

	c.GET("", handler.FetchG)
}

func (eGH *examGroupHandler) FetchG(eC echo.Context) error {
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

	if courseIDP, ok := eC.QueryParams()["course"]; ok {
		courseIDHex, err := primitive.ObjectIDFromHex(courseIDP[0])
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		filter.CourseID = courseIDHex
	}

	if classIDP, ok := eC.QueryParams()["class"]; ok {
		classIDHex, err := primitive.ObjectIDFromHex(classIDP[0])
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		filter.ClassID = classIDHex
	}

	if levelIDP, ok := eC.QueryParams()["level"]; ok {
		levelIDHex, err := primitive.ObjectIDFromHex(levelIDP[0])
		if err != nil {
			return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		}
		filter.LevelID = levelIDHex
	}

	courseGs, err := eGH.eGUsecase.FetchG(filter)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, courseGs)
}
