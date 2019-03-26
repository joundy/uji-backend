package http

import (
	"net/http"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/auth"

	"github.com/labstack/echo"
)

//AuthHandler represent handler for auth
type AuthHandler struct {
	aUsecase auth.Usecase
}

//NewAuthHandler represent initialitation authhandler
func NewAuthHandler(c *echo.Echo, aU auth.Usecase) {
	handler := &AuthHandler{aU}
	c.POST("/auth", handler.Auth)
}

type credential struct {
	Email    string
	Password string
}

//Auth is method from Authhandler
func (aH *AuthHandler) Auth(eC echo.Context) error {

	var ct credential
	eC.Bind(&ct)

	//usecase
	auth, err := aH.aUsecase.Auth(ct.Email, ct.Password)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, auth)
}
