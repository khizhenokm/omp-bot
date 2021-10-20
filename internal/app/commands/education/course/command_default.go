package course

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CourseCommander) Default(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Unsupported command: "+inputMessage.Text)
	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
