package http

import (
	"net/http"

	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/usecase/file"

	"github.com/labstack/echo"
)

//FileHandler represent handler for fileHandler
type fileHandler struct {
	fUsecase file.Usecase
}

//NewFileHandler represent initialitation fileHandler
func NewFileHandler(e *echo.Echo, fU file.Usecase) {
	handler := &fileHandler{fU}
	c := e.Group("/files")

	c.POST("", handler.Upload)
	c.GET("/:filename", handler.Stream)
}

func (fH *fileHandler) Stream(eC echo.Context) error {
	filename := eC.Param("filename")

	file, err := fH.fUsecase.Stream(&filename)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.Stream(http.StatusOK, "image/png", file)
}

func (fH *fileHandler) Upload(eC echo.Context) error {

	file, err := eC.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filename := file.Filename

	result, err := fH.fUsecase.Upload(&filename, &src)
	if err != nil {
		return eC.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return eC.JSON(http.StatusOK, result)
}
