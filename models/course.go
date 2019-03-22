package models

import "time"

//Course is represent model for course data
type Course struct {
	Title       string
	Description string
	Deadline    time.Time
	ExamGroupID string
}

//CourseG is represent model for courseGroup data
type CourseG struct {
	data  []Course
	count int
}
