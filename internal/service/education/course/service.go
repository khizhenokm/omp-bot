package course

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/education"
)

type DummyCourseService struct {
	courses    []education.Course
	idSequence uint64
}

func NewDummyCourseService() *DummyCourseService {
	return &DummyCourseService{
		courses: []education.Course{
			{
				Id:          1,
				Title:       "One",
				Description: "One description",
			},
		},
		idSequence: 1,
	}
}

func (s *DummyCourseService) Describe(courseID uint64) (*education.Course, error) {
	courseIndex := s.FindElementIndex(courseID)
	if courseIndex == -1 {
		err := fmt.Errorf("course with id %d not found", courseID)
		return nil, err
	}

	return &s.courses[courseIndex], nil
}

func (s *DummyCourseService) List(cursor uint64, limit uint64) ([]education.Course, error) {
	if uint64(len(s.courses)) < cursor {
		return nil, fmt.Errorf("the cursor is out of range")
	}

	if cursor+limit > uint64(len(s.courses)) {
		return s.courses[cursor:len(s.courses)], nil
	}

	return s.courses[cursor : cursor+limit], nil
}

func (s *DummyCourseService) Create(course education.Course) (uint64, error) {
	s.idSequence++
	course.Id = s.idSequence
	s.courses = append(s.courses, course)
	return course.Id, nil
}

func (s *DummyCourseService) Update(courseID uint64, course education.Course) error {
	courseIndex := s.FindElementIndex(courseID)
	if courseIndex == -1 {
		err := fmt.Errorf("course with id %d not found", courseID)
		return err
	}

	s.courses[courseIndex] = course
	return nil
}

func (s *DummyCourseService) Remove(courseID uint64) (bool, error) {
	courseIndex := s.FindElementIndex(courseID)
	if courseIndex == -1 {
		err := fmt.Errorf("course with id %d not found", courseID)
		return false, err
	}

	s.courses = append(s.courses[:courseIndex], s.courses[courseIndex+1:]...)
	return true, nil
}

func (c *DummyCourseService) FindElementIndex(courseID uint64) int {
	if c.idSequence < courseID {
		return -1
	}

	for index, course := range c.courses {
		if course.Id == courseID {
			return index
		}
	}

	return -1
}
