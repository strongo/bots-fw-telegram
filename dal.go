package telegram

import (
	"context"
	"github.com/bots-go-framework/bots-fw-telegram-models/botsfwtgmodels"
)

type DataStore interface {
	TgChatInstanceDal
}

// TgChatInstanceDal is Data Access Layer for telegram chat instance Data
type TgChatInstanceDal interface {
	GetTelegramChatInstanceByID(c context.Context, id string) (tgChatInstance botsfwtgmodels.TgChatInstanceData, err error)
	NewTelegramChatInstance(chatInstanceID string, chatID int64, preferredLanguage string) (tgChatInstance botsfwtgmodels.TgChatInstanceData)
	SaveTelegramChatInstance(c context.Context, tgChatInstance botsfwtgmodels.TgChatInstanceData) (err error)
}

var tgChatInstanceDal TgChatInstanceDal
