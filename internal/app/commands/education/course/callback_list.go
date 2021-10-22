package course

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

func NewCallbackListData(offset uint64, limit uint64) *CallbackListData {
	return &CallbackListData{
		Offset: offset,
		Limit:  limit,
	}
}

func (c *CourseCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) error {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		return err
	}

	err = c.ListInternal(callback.Message.Chat.ID, parsedData.Offset, parsedData.Limit)
	if err != nil {
		return err
	}

	return nil
}

func (d CallbackListData) ToJsonString() (string, error) {
	serializedData, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(serializedData), nil
}
