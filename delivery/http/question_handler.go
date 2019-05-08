package http

import (
	"fmt"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/question"
	"github.com/labstack/echo"
)

//examGroupHandler represent handler for course
type questionHandler struct {
	qUsecase question.Usecase
}

//NewQuestionHandler represent initialization NewQuestionHandler
func NewQuestionHandler(e *echo.Echo, qU question.Usecase) {
	handler := &questionHandler{qU}

	c := e.Group("/questions")

	c.POST("", handler.Create)
}

func (eH *questionHandler) Create(eC echo.Context) error {

	var question models.Question
	eC.Bind(&question)

	fmt.Println(question)

	// question.ID = primitive.NewObjectID()

	// if examGroupIDP, ok := eC.QueryParams()["examGroupId"]; ok {
	// 	examGroupIDHex, err := primitive.ObjectIDFromHex(examGroupIDP[0])
	// 	if err != nil {
	// 		return eC.JSON(http.StatusBadRequest, models.ResponseError{Message: err.Error()})
	// 	}
	// 	exam.ExamGroupID = examGroupIDHex
	// }

	// resID, err := eH.eUsecase.Create(&exam)
	// if err != nil {
	// 	return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	// }

	// return eC.JSON(http.StatusOK, resID)
	return nil
}
