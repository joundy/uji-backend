package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/haffjjj/uji-backend/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spf13/viper"

	_httpDelivery "github.com/haffjjj/uji-backend/delivery/http"

	_courseRepo "github.com/haffjjj/uji-backend/repository/course"
	_examRepo "github.com/haffjjj/uji-backend/repository/exam"
	_examGroupRepo "github.com/haffjjj/uji-backend/repository/examgroup"
	_examLogRepo "github.com/haffjjj/uji-backend/repository/examlog"
	_questionRepo "github.com/haffjjj/uji-backend/repository/question"
	_userRepo "github.com/haffjjj/uji-backend/repository/user"

	_authUsecase "github.com/haffjjj/uji-backend/usecase/auth"
	_courseUsecase "github.com/haffjjj/uji-backend/usecase/course"
	_examUsecase "github.com/haffjjj/uji-backend/usecase/exam"
	_examGroupUsecase "github.com/haffjjj/uji-backend/usecase/examgroup"
	_examLogUsecase "github.com/haffjjj/uji-backend/usecase/examlog"
)

func init() {
	viper.SetConfigFile("config")
	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		dbHost  = viper.GetString("database.mongodb.host")
		dbPort  = viper.GetString("database.mongodb.port")
		dbUName = viper.GetString("database.mongodb.username")
		dbPass  = viper.GetString("database.mongodb.password")
		port    = viper.GetString("port")
	)

	mgoClient, err := mongo.Connect(context.TODO(), fmt.Sprint("mongodb://", dbUName, ":", dbPass, "@", dbHost, dbPort))
	if err != nil {
		log.Fatal(err)
	}
	err = mgoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer mgoClient.Disconnect(context.TODO())

	// ===========

	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = &utils.Validator{Validator: validator.New()}

	// ===========

	courseRepo := _courseRepo.NewMongoCourseRepository(mgoClient)
	examGroupRepo := _examGroupRepo.NewMongoExamGroupRepository(mgoClient)
	examRepo := _examRepo.NewMongoExamRepository(mgoClient)
	examLogRepo := _examLogRepo.NewMongoExamLogRepository(mgoClient)
	questionRepo := _questionRepo.NewMongoQuestionRepository(mgoClient)
	userRepo := _userRepo.NewMongoUserRespository(mgoClient)

	courseUsecase := _courseUsecase.NewCourseUsecase(courseRepo)
	examGroupUsecase := _examGroupUsecase.NewExamGroupUsecase(examGroupRepo)
	examUsecase := _examUsecase.NewExamUsecase(examRepo)
	examLogUsecase := _examLogUsecase.NewExamLogUsecase(examLogRepo, examRepo, questionRepo)
	authUsecase := _authUsecase.NewAuthUsecase(userRepo)

	_httpDelivery.NewCourseHandler(e, courseUsecase)
	_httpDelivery.NewExamGroupHandler(e, examGroupUsecase)
	_httpDelivery.NewExamHandler(e, examUsecase)
	_httpDelivery.NewExamLogHandler(e, examLogUsecase)
	_httpDelivery.NewAuthHandler(e, authUsecase)

	// ===========

	e.Logger.Fatal(e.Start(port))
}
