package course

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

func (c *CourseCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var msg tgbotapi.MessageConfig
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		msg = tgbotapi.NewMessage(callback.Message.Chat.ID, "Wrong argument format")
	} else {
		courses, err := c.courseService.List(parsedData.Offset, parsedData.Limit)
		if err != nil {
			log.Printf("CourseCommander.List: error getting courses - %v", err)
		}

		outputMsgText := "Here all the courses: \n\n"
		for _, course := range courses {
			outputMsgText += course.String()
			outputMsgText += "\n"
		}

		msg = tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
		var buttons []tgbotapi.InlineKeyboardButton

		if parsedData.Offset > 0 {
			data := CallbackListData{
				Offset: parsedData.Offset - parsedData.Limit,
				Limit:  parsedData.Limit,
			}
			serializedData, _ := json.Marshal(data)
			callbackPath := path.CallbackPath{
				Domain:       "education",
				Subdomain:    "course",
				CallbackName: "list",
				CallbackData: string(serializedData),
			}
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPath.String()))
		}

		if uint64(len(courses)) == parsedData.Limit {
			data := CallbackListData{
				Offset: parsedData.Offset + parsedData.Limit,
				Limit:  parsedData.Limit,
			}
			serializedData, _ := json.Marshal(data)
			callbackPath := path.CallbackPath{
				Domain:       "education",
				Subdomain:    "course",
				CallbackName: "list",
				CallbackData: string(serializedData),
			}
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()))
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
