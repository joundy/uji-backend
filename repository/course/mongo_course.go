package course

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoCourseRepository struct {
	mgoClient *mongo.Client
}

//NewMongoCourseRepository represent initialization mongoCourseRepository
func NewMongoCourseRepository(c *mongo.Client) Repository {
	return &mongoCourseRepository{c}
}

func (m *mongoCourseRepository) FetchG() ([]*models.CourseG, error) {
	return nil, nil
}
