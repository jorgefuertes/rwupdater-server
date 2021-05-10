package cfg

import (
	"time"
)

// AppName - Applicaton name
const AppName = "RetroWiki FPGA Files Server"

// Author - Application author
const Author = "©2021 Jorge Fuertes & Ramón Martinez"

// smtp
var SMTPHost = "undefined"
var SMTPPort = "undefined"
var SMTPUser = "undefined"
var SMTPPass = "undefined"
var SMTPTo = "undefined"

// MainConfig - main config type
type MainConfig struct {
	Env    *string
	Server struct {
		IP          *string
		Port        *string
		BodyLimitMb *int
		RTimeout    *time.Duration
		WTimeout    *time.Duration
		Concurrency *int
	}
	Root string
}

// AvailableLangRef - Static regexp
const AvailableLangReg = `^(en|es)$`

// Main - Main configuration
var Main MainConfig

// IsDev - Boolean
func IsDev() bool {
	return *Main.Env == "dev"
}

// IsTest - Boolean
func IsTest() bool {
	return *Main.Env == "test"
}

// IsProd - Boolean
func IsProd() bool {
	return *Main.Env == "prod"
}
