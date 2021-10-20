package course

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CourseCommander) Delete(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()
	courseId, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return NewBadRequestError("Wrong course Id format")
	}

	_, err = c.courseService.Remove(courseId)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Course succesfully deleted")
	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
