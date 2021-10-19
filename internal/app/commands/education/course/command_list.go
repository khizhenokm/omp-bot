package course

import (
	"encoding/json"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CourseCommander) List(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	limit, err := strconv.ParseUint(args, 10, 64)
	if err != nil || limit == 0 {
		limit = 5
	}

	courses, err := c.courseService.List(0, limit)
	if err != nil {
		log.Printf("CourseCommander.List: error getting courses - %v", err)
	}

	outputMsgText := "Here all the courses: \n\n"
	for _, course := range courses {
		outputMsgText += course.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if uint64(len(courses)) == limit {

		data := CallbackListData{
			Offset: limit,
			Limit:  limit,
		}

		serializedData, _ := json.Marshal(data)
		callbackPath := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "course",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseCommander.List: error sending reply message to chat - %v", err)
	}
}
