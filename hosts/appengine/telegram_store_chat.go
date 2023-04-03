package gaehost

import (
	//"fmt"
	"context"
	"github.com/strongo/bots-framework/botsfw"
	"github.com/strongo/bots-fw-dalgo/dalgo4botsfw"
	telegram "github.com/strongo/bots-fw-telegram"
	gae "github.com/strongo/bots-host-gae"
	"github.com/strongo/nds"
	"google.golang.org/appengine/datastore"
	"strconv"
	"time"
	//"reflect"
)

// GaeTelegramChatStore DAL to telegram chat entity
type GaeTelegramChatStore struct {
	gae.GaeBotChatStore
}

//var _ botsfw.BotChatStore = (*GaeTelegramChatStore)(nil) // Check for interface implementation at compile time

// §NewGaeTelegramChatStore creates DAL to Telegram chat entity
func NewGaeTelegramChatStore(newTelegramChatData func() botsfw.BotChat) botsfw.BotChatStore {
	{ // Validate newTelegramChatData
		//tgChatData := newTelegramChatData()
		//if _, ok := tgChatData.(*telegram.TgChatEntityBase); !ok {
		//	v := reflect.ValueOf(tgChatData)
		//	if v.Type() != reflect.TypeOf(telegram.TgChatEntityBase{}) {
		//		panic(fmt.Sprintf("Expected *telegram.TelegramChat but received %T", entity))
		//	}
		//}
	}
	newTelegramChatData()
	//_ = &GaeTelegramChatStore{
	//	GaeBotChatStore: *gae.NewGaeBotChatStore(telegram.ChatKind, nil, nil, newTelegramChatEntity),
	//}
	return dalgo4botsfw.NewBotChatStore(telegram.ChatKind, nil, func() botsfw.BotChat {
		tgChatData := newTelegramChatData()
		return tgChatData
	})
}

// MarkTelegramChatAsForbidden marks tg chat as forbidden
func MarkTelegramChatAsForbidden(c context.Context, botID string, tgChatID int64, dtForbidden time.Time) error {
	return nds.RunInTransaction(c, func(c context.Context) (err error) {
		key := datastore.NewKey(c, telegram.ChatKind, botsfw.NewChatID(botID, strconv.FormatInt(tgChatID, 10)), 0, nil)
		var chat telegram.TgChatEntityBase
		if err = nds.Get(c, key, &chat); err != nil {
			return
		}
		var changed bool
		if chat.DtForbidden.IsZero() {
			chat.DtForbidden = dtForbidden
			changed = true
		}

		if chat.DtForbiddenLast.IsZero() || chat.DtForbiddenLast.Before(dtForbidden) {
			chat.DtForbiddenLast = dtForbidden
			changed = true
		}

		if changed {
			_, err = nds.Put(c, key, &chat)
		}
		return
	}, nil)
}
