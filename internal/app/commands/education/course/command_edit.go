package course

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

func (c *CourseCommander) Edit(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()
	deserializedData := education.Course{}
	err := json.Unmarshal([]byte(args), &deserializedData)
	if err != nil {
		return NewBadRequestError("Wrong arguments format")
	}

	err = c.courseService.Update(deserializedData.Id, deserializedData)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Course succesfully updated")
	if _, err := c.bot.Send(msg); err != nil {
		return err
	}

	return nil
}
