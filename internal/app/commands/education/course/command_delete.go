package course

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CourseCommander) Delete(inputMessage *tgbotapi.Message) {

	var msg tgbotapi.MessageConfig

	args := inputMessage.CommandArguments()
	courseId, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong Course Id format")
	} else {
		_, err := c.courseService.Remove(courseId)
		if err != nil {
			//TODO: Map to userfriendly message
			log.Printf("fail to remove course with id %d: %v", courseId, err)
			return
		}

		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "Course succesfully deleted")
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseCommander.Delete: error sending reply message to chat - %v", err)
	}
}
