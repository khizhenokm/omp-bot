package course

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

func (c *CourseCommander) New(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()
	deserializedData := education.Course{}
	err := json.Unmarshal([]byte(args), &deserializedData)
	if err != nil {
		return NewBadRequestError("Wrong arguments format")
	}

	courseId, err := c.courseService.Create(deserializedData)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Course created with id %d", courseId))

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
