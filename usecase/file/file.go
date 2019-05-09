package file

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/haffjjj/uji-backend/models"
)

type fileUsecase struct{}

//NewFileUsecase represent initializatin NewQuestionUsecase
func NewFileUsecase() Usecase {
	return &fileUsecase{}
}

func (fU *fileUsecase) Stream(filename *string) (*os.File, error) {
	file, err := os.Open("images/" + *filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fU *fileUsecase) Upload(filename *string, src *multipart.File) (*models.File, error) {
	dst, err := os.Create("images/" + *filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, *src); err != nil {
		return nil, err
	}

	result := models.File{
		Filename: *filename,
	}

	return &result, nil
}
