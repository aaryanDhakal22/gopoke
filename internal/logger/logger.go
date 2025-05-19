package logger

import (
	"log/slog"
	"os"
)

var logLevel = new(slog.LevelVar)
var Log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: logLevel,
}))

func Init() {
	slog.SetDefault(Log)
}

func SetLevel(level slog.Level) {
	logLevel.Set(level)
}
