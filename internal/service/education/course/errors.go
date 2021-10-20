package course

import (
	"fmt"
)

type CourseNotFoundError struct {
	courseId uint64
}

func NewCourseNotFoundError(id uint64) *CourseNotFoundError {
	return &CourseNotFoundError{
		courseId: id,
	}
}

func (e *CourseNotFoundError) Error() string {
	return fmt.Sprintf("course with id %d not found", e.courseId)
}
