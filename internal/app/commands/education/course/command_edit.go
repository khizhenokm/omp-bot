package course

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

func (c *CourseCommander) Edit(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	deserializedData := education.Course{}
	err := json.Unmarshal([]byte(args), &deserializedData)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	err = c.courseService.Update(deserializedData.Id, deserializedData)
	if err != nil {
		log.Printf("fail to get update course: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Succesfully updated")

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseCommander.Edit: error sending reply message to chat - %v", err)
	}
}
