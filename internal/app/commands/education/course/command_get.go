package course

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CourseCommander) Get(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()
	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return NewBadRequestError("Wrong course Id format")
	}

	course, err := c.courseService.Describe(idx)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, course.String())

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
