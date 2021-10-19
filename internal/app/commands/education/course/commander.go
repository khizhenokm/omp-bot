package course

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"github.com/ozonmp/omp-bot/internal/service/education/course"
)

type CourseService interface {
	Describe(courseID uint64) (*education.Course, error)
	List(cursor uint64, limit uint64) ([]education.Course, error)
	Create(course education.Course) (uint64, error)
	Update(courseID uint64, course education.Course) error
	Remove(courseID uint64) (bool, error)
}

type CourseCommander struct {
	bot           *tgbotapi.BotAPI
	courseService CourseService
}

func NewEducationCourseCommander(bot *tgbotapi.BotAPI) *CourseCommander {
	//TODO: Move service creation to composition root
	service := course.NewDummyCourseService()
	return &CourseCommander{
		bot:           bot,
		courseService: service,
	}
}

func (c *CourseCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CourseCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CourseCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "new":
		c.New(msg)
	case "get":
		c.Get(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
