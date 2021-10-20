package course

import (
	"errors"

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
	courseIndex, err := s.GetElementIndex(courseID)
	if err != nil {
		return nil, err
	}

	return &s.courses[courseIndex], nil
}

func (s *DummyCourseService) List(cursor uint64, limit uint64) ([]education.Course, error) {
	if uint64(len(s.courses)) < cursor {
		return nil, errors.New("the cursor is out of range")
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
	if courseID != course.Id {
		err := errors.New("something goes wrong")
		return err
	}

	courseIndex, err := s.GetElementIndex(courseID)
	if err != nil {
		return err
	}

	s.courses[courseIndex] = course
	return nil
}

func (s *DummyCourseService) Remove(courseID uint64) (bool, error) {
	courseIndex, err := s.GetElementIndex(courseID)
	if err != nil {
		return false, err
	}

	s.courses = append(s.courses[:courseIndex], s.courses[courseIndex+1:]...)
	return true, nil
}

func (c *DummyCourseService) Count() int {
	return len(c.courses)
}

func (c *DummyCourseService) GetElementIndex(courseID uint64) (int, error) {
	for index, course := range c.courses {
		if course.Id == courseID {
			return index, nil
		}
	}

	return 0, NewCourseNotFoundError(courseID)
}
