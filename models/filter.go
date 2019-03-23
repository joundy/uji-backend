package models

//Filter represent model for Filter data
type Filter struct {
	Start       int
	Limit       int
	ClassID     string
	LevelID     string
	ExamGroupID string
	CourseID    string
}
