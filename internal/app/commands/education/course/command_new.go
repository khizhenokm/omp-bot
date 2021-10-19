package course

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

func (c *CourseCommander) New(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	deserializedData := education.Course{}
	err := json.Unmarshal([]byte(args), &deserializedData)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	courseId, err := c.courseService.Create(deserializedData)
	if err != nil {
		log.Printf("fail to create course: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Created with id %d", courseId),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseCommander.New: error sending reply message to chat - %v", err)
	}
}
