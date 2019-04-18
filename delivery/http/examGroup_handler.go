package http

import (
	"net/http"
	"strconv"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/examgroup"
	"github.com/labstack/echo"
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

	if classP, ok := eC.QueryParams()["class"]; ok {
		mF.Class = classP[0]
	}

	if levelP, ok := eC.QueryParams()["level"]; ok {
		mF.Level = levelP[0]
	}

	if tagP, ok := eC.QueryParams()["tag"]; ok {
		mF.Tag = tagP[0]
	}

	//usecase
	courseGs, err := eGH.eGUsecase.FetchG(&mF)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, courseGs)
}
