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
	c.POST("", handler.Create)
	c.GET("/:id", handler.GetByID)
	c.PUT("/:id", handler.UpdateByID)
	c.DELETE("/:id", handler.DeleteByID)
}

func (eGH *examGroupHandler) GetByID(eC echo.Context) error {
	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	examGroup, err := eGH.eGUsecase.GetByID(&ID)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examGroup)
}

func (eGH *examGroupHandler) UpdateByID(eC echo.Context) error {
	var examGroup models.ExamGroup
	eC.Bind(&examGroup)

	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	err = eGH.eGUsecase.UpdateByID(&ID, &examGroup)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (eGH *examGroupHandler) DeleteByID(eC echo.Context) error {
	IDP := eC.Param("id")
	ID, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	err = eGH.eGUsecase.DeleteByID(&ID)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (eGH *examGroupHandler) Create(eC echo.Context) error {
	var examGroup models.ExamGroup
	eC.Bind(&examGroup)

	examGroup.ID = primitive.NewObjectID()

	resID, err := eGH.eGUsecase.Create(&examGroup)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, resID)
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
