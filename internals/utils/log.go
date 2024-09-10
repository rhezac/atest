package utils

import (
	"os"

	"golang.org/x/exp/slog"
)

type ServiceLogInfo struct {
	HttpMethod string
	UrlPath    string
	HttpCode   string
	HttpStatus string
}

func InitLog(config *Config) {
	logFile, err := os.OpenFile(config.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	handler := slog.NewJSONHandler(logFile, opts)

	logger := slog.New(handler)
	slog.SetDefault(logger)

}
