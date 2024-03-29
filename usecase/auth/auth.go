package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/user"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	uRepository user.Repository
}

//NewAuthUsecase represent initializatin courseUsecase
func NewAuthUsecase(uR user.Repository) Usecase {
	return &authUsecase{uR}
}

func (c *authUsecase) AuthGuest() (*models.ResToken, error) {

	userID := primitive.NewObjectID()
	fullName := "Guest"
	email := "hafizjoundys@gmail.com"

	//hardcode id from databases
	userTypeID, err := primitive.ObjectIDFromHex("5caf9f4dc160d14f26dc2929")
	if err != nil {
		return nil, err
	}

	claims := &models.JWTClaims{
		userID,
		fullName,
		email,
		userTypeID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	JWTsecret := viper.GetString("jwtSecret")

	tSigned, err := token.SignedString([]byte(JWTsecret))
	if err != nil {
		return nil, err
	}

	return &models.ResToken{tSigned}, nil
}

func (c *authUsecase) Auth(email, password string) (*models.ResToken, error) {

	user, err := c.uRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	claims := &models.JWTClaims{
		user.ID,
		user.Fullname,
		user.Email,
		user.UserTypeID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	JWTsecret := viper.GetString("jwtSecret")

	tSigned, err := token.SignedString([]byte(JWTsecret))
	if err != nil {
		return nil, err
	}

	return &models.ResToken{tSigned}, nil
}
