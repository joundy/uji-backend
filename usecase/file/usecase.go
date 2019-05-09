package file

import (
	"mime/multipart"
	"os"

	"github.com/haffjjj/uji-backend/models"
)

//Usecase represent course usecase contract
type Usecase interface {
	Upload(filename *string, src *multipart.File) (*models.File, error)
	Stream(filename *string) (*os.File, error)
}
