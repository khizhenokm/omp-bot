package course

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CourseCommander) List(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()
	limit, err := strconv.ParseUint(args, 10, 64)
	if err != nil || limit == 0 {
		limit = 5
	}

	err = c.ListInternal(inputMessage.Chat.ID, 0, limit)
	if err != nil {
		return err
	}

	return nil
}

func (c *CourseCommander) ListInternal(chatId int64, offset uint64, limit uint64) error {
	totalNumberOfCourses := c.courseService.Count()
	courses, err := c.courseService.List(offset, limit)
	if err != nil {
		return err
	}

	outputMsgText := "Here all the courses: \n\n"
	for _, course := range courses {
		outputMsgText += course.String()
		outputMsgText += "\n"
	}

	var buttons []tgbotapi.InlineKeyboardButton
	if offset > 0 {
		button, err := CreateCallbackListButton("Previous page", offset-limit, limit)
		if err != nil {
			return err
		}
		buttons = append(buttons, *button)
	}

	if offset+limit < uint64(totalNumberOfCourses) {
		button, err := CreateCallbackListButton("Next page", offset+limit, limit)
		if err != nil {
			return err
		}
		buttons = append(buttons, *button)
	}

	msg := tgbotapi.NewMessage(chatId, outputMsgText)

	if buttons != nil {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func CreateCallbackListButton(buttonName string, offset uint64, limit uint64) (*tgbotapi.InlineKeyboardButton, error) {
	data, err := NewCallbackListData(offset, limit).ToJsonString()
	if err != nil {
		return nil, err
	}

	callbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "course",
		CallbackName: "list",
		CallbackData: data,
	}

	button := tgbotapi.NewInlineKeyboardButtonData(buttonName, callbackPath.String())

	return &button, nil
}
