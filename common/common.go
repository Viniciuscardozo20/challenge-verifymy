package common

import "github.com/gofiber/fiber/v2/log"

type Data struct {
	Status  string `status:"code"`
	Message string `json:"message"`
}

func ParseLogLevel(level string) log.Level {
	switch level {
	case "info":
		return log.LevelInfo
	case "error":
		return log.LevelError
	case "debug":
		return log.LevelDebug
	default:
		return log.LevelInfo
	}
}
