package course

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CourseCommander) Get(inputMessage *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong course Id format")
	} else {
		course, err := c.courseService.Describe(idx)
		if err != nil {
			//TODO: Map to userfriendly message
			log.Printf("fail to get course with idx %d: %v", idx, err)
			return
		}

		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, course.String())
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseCommander.Get: error sending reply message to chat - %v", err)
	}
}
