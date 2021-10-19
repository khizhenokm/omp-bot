package course

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CourseCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__education__course — print list of commands\n"+
			"/get__education__course {courseId} — get a course\n"+
			"/list__education__course {limit} — get courses\n"+
			"/delete__education__course {courseId} — delete a course\n"+
			"/new__education__course {\"title\": \"This is a title\", \"description\": \"This is an description\"} — create a course\n"+
			"/edit__education__course {\"id\": 1234,\"title\": \"This is a title\", \"description\": \"This is an description\"} — edit a course\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseCommander.Help: error sending reply message to chat - %v", err)
	}
}
