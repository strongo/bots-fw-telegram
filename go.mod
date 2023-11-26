module github.com/bots-go-framework/bots-fw-telegram

go 1.21

toolchain go1.21.4

//replace github.com/bots-go-framework/bots-fw-store => ../bots-fw-store

//replace github.com/bots-go-framework/bots-fw => ../bots-fw

require (
	github.com/bots-go-framework/bots-api-telegram v0.4.2
	github.com/bots-go-framework/bots-fw v0.23.3
	github.com/bots-go-framework/bots-fw-store v0.4.0
	github.com/bots-go-framework/bots-fw-telegram-models v0.0.15
	github.com/dal-go/dalgo v0.12.0
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7
	github.com/strongo/i18n v0.0.4
	github.com/strongo/log v0.3.0
	github.com/strongo/strongoapp v0.9.1
)

require (
	github.com/bots-go-framework/bots-go-core v0.0.2 // indirect
	github.com/strongo/gamp v0.0.1 // indirect
	github.com/strongo/random v0.0.1 // indirect
	github.com/strongo/validation v0.0.6 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
)
