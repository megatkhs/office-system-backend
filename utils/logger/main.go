package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

var logger *slog.Logger

func Initialize() {
	fileName := fmt.Sprintf("logs/log_%s.log", time.Now().Format("20060102"))
	logFile, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	multiLogFile := io.MultiWriter(os.Stdout, logFile)

	h := slog.NewJSONHandler(multiLogFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger = slog.New(h)
}

func Debug(msg string, fields ...slog.Attr) {
	logger.Debug(msg, "payload", slog.GroupValue(fields...))
}

func Info(msg string, fields ...slog.Attr) {
	logger.Info(msg, "payload", slog.GroupValue(fields...))
}

func Warn(msg string, fields ...slog.Attr) {
	logger.Warn(msg, "payload", slog.GroupValue(fields...))
}

func Error(msg string, err error) {
	logger.Error(msg, err)
}

func Fatal(msg string, err error) {
	logger.Error(msg, err)
	os.Exit(1)
}
