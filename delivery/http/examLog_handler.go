package http

import (
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
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
	m := &Middleware{}

	gAuth := e.Group("/examLogs")
	g := e.Group("/examLogs")
	gAuth.Use(m.JWTAuth())

	g.POST("/generate/guest", handler.GenerateGuest)
	gAuth.POST("/generate", handler.Generate)

	gAuth.GET("/:id", handler.GetByIDAndStart)
	g.GET("/:id/guest", handler.GetByIDAndStartGuest)
	gAuth.GET("", handler.FetchG)
	gAuth.PUT("/:id/setAnswers/questions/:questionId", handler.SetAnswer)
	g.PUT("/:id/guest/setAnswers/questions/:questionId", handler.SetAnswerGuest)
	gAuth.POST("/:id/submit", handler.Submit)
	g.POST("/:id/guest/submit", handler.Submit)
}

func (eLH *examLogHandler) FetchG(eC echo.Context) error {
	user := eC.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

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

	userIDHex, err := primitive.ObjectIDFromHex(claims["ID"].(string))
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	//usecase
	examLogGs, err := eLH.eGUsecase.FetchG(&userIDHex, &mF)
	if err != nil {
		eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examLogGs)
}

func (eLH *examLogHandler) Generate(eC echo.Context) error {
	user := eC.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	examIDF := eC.FormValue("examId")
	examIDHex, err := primitive.ObjectIDFromHex(examIDF)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex(claims["ID"].(string))
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	//usecase
	resID, err := eLH.eGUsecase.Generate(userIDHex, examIDHex, false)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, resID)
}

func (eLH *examLogHandler) GenerateGuest(eC echo.Context) error {
	examIDF := eC.FormValue("examId")
	examIDHex, err := primitive.ObjectIDFromHex(examIDF)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	var userIDHex primitive.ObjectID

	//usecase
	resID, err := eLH.eGUsecase.Generate(userIDHex, examIDHex, true)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, resID)
}

func (eLH *examLogHandler) GetByIDAndStart(eC echo.Context) error {
	user := eC.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex(claims["ID"].(string))
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	//usecase
	examLog, err := eLH.eGUsecase.GetByIDAndStart(&IDHex, &userIDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examLog)
}

func (eLH *examLogHandler) GetByIDAndStartGuest(eC echo.Context) error {
	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	var userIDHex primitive.ObjectID

	//usecase
	examLog, err := eLH.eGUsecase.GetByIDAndStart(&IDHex, &userIDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, examLog)
}

func (eLH *examLogHandler) SetAnswer(eC echo.Context) error {
	user := eC.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	questionIDP := eC.Param("questionId")
	questionIDHex, err := primitive.ObjectIDFromHex(questionIDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex(claims["ID"].(string))
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	var isSelectedIdsHex []primitive.ObjectID

	fParams, _ := eC.FormParams()
	isSelectedIds := fParams["isSelectedId"]

	for _, v := range isSelectedIds {
		elemHex, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		}

		isSelectedIdsHex = append(isSelectedIdsHex, elemHex)
	}

	//usecase
	err = eLH.eGUsecase.SetAnswer(&IDHex, &userIDHex, &questionIDHex, &isSelectedIdsHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (eLH *examLogHandler) SetAnswerGuest(eC echo.Context) error {
	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	questionIDP := eC.Param("questionId")
	questionIDHex, err := primitive.ObjectIDFromHex(questionIDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	var userIDHex primitive.ObjectID

	var isSelectedIdsHex []primitive.ObjectID

	fParams, _ := eC.FormParams()
	isSelectedIds := fParams["isSelectedId"]

	for _, v := range isSelectedIds {
		elemHex, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		}

		isSelectedIdsHex = append(isSelectedIdsHex, elemHex)
	}

	//usecase
	err = eLH.eGUsecase.SetAnswer(&IDHex, &userIDHex, &questionIDHex, &isSelectedIdsHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (eLH *examLogHandler) Submit(eC echo.Context) error {
	user := eC.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	userIDHex, err := primitive.ObjectIDFromHex(claims["ID"].(string))
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	//usecase
	err = eLH.eGUsecase.Submit(&IDHex, &userIDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}

func (eLH *examLogHandler) SubmitGuest(eC echo.Context) error {
	IDP := eC.Param("id")
	IDHex, err := primitive.ObjectIDFromHex(IDP)
	if err != nil {
		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	}

	var userIDHex primitive.ObjectID

	//usecase
	err = eLH.eGUsecase.Submit(&IDHex, &userIDHex)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusNoContent, "")
}
