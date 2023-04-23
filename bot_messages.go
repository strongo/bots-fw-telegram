package telegram

import (
	"github.com/bots-go-framework/bots-api-telegram/tgbotapi"
	"github.com/bots-go-framework/bots-fw/botsfw"
)

// InlineBotMessage is wrapper for Telegram bot message
type InlineBotMessage struct {
	*tgbotapi.InlineConfig
}

// BotMessageType returns BotMessageTypeInlineResults
func (InlineBotMessage) BotMessageType() botsfw.BotMessageType {
	return botsfw.BotMessageTypeInlineResults
}

// CallbackAnswer is callback answer message
type CallbackAnswer tgbotapi.AnswerCallbackQueryConfig

// BotMessageType returns BotMessageTypeCallbackAnswer
func (CallbackAnswer) BotMessageType() botsfw.BotMessageType {
	return botsfw.BotMessageTypeCallbackAnswer
}

// LeaveChat is leave chat message from bot
type LeaveChat tgbotapi.LeaveChatConfig

// BotMessageType return BotMessageTypeLeaveChat
func (LeaveChat) BotMessageType() botsfw.BotMessageType {
	return botsfw.BotMessageTypeLeaveChat
}

// ExportChatInviteLink is TG message
type ExportChatInviteLink tgbotapi.ExportChatInviteLink

// BotMessageType returns BotMessageTypeExportChatInviteLink
func (ExportChatInviteLink) BotMessageType() botsfw.BotMessageType {
	return botsfw.BotMessageTypeExportChatInviteLink
}
