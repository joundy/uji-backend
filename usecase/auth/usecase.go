package auth

import "github.com/haffjjj/uji-backend/models"

//Usecase represent course usecase contract
type Usecase interface {
	Auth(email, password string) (*models.ResToken, error)
	AuthGuest() (*models.ResToken, error)
}
