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
	Count() int
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
	var err error
	switch callbackPath.CallbackName {
	case "list":
		err = c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CourseCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}

	if err != nil {
		c.HandleError(callback.Message.Chat.ID, err)
	}
}

func (c *CourseCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error
	switch commandPath.CommandName {
	case "help":
		err = c.Help(msg)
	case "list":
		err = c.List(msg)
	case "new":
		err = c.New(msg)
	case "get":
		err = c.Get(msg)
	case "edit":
		err = c.Edit(msg)
	case "delete":
		err = c.Delete(msg)
	default:
		err = c.Default(msg)
	}

	if err != nil {
		c.HandleError(msg.Chat.ID, err)
	}
}

func (c *CourseCommander) HandleError(chatId int64, err error) {
	var msg tgbotapi.MessageConfig
	switch err.(type) {
	case *course.CourseNotFoundError, *BadRequestError:
		msg = tgbotapi.NewMessage(chatId, err.Error())
	default:
		log.Printf("CourseCommander.HandleError: Unknow error happend - %v", err)
		msg = tgbotapi.NewMessage(chatId, "Unknow error happend pls contact bot administrator")
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseCommander.HandleError: error sending reply message to chat - %v", err)
	}
}
