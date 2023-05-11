package telegram

import "github.com/bots-go-framework/bots-fw/botsfw"

// Platform is a bots platform descriptor (in this case - for Telegram)
var Platform botsfw.BotPlatform = platform{}

// platform describes Telegram platform
type platform struct {
}

// PlatformID is 'telegram'
const PlatformID = "telegram"

// ID returns 'telegram'
func (p platform) ID() string {
	return PlatformID
}

// Version returns '2.0'
func (p platform) Version() string {
	return "2.0"
}
