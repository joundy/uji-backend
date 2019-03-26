package user

import "github.com/haffjjj/myblog-backend/models"

//Repository represent course repository contract
type Repository interface {
	GetByEmail(email string) (*models.User, error)
}
