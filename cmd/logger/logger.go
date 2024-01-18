package logger

import (
	"log/slog"
	"os"
)

// cmd это точка входа для приложения, убрать весь логер отсюда например в папку packeges
func SetupLogger(env string) *slog.Logger {
	log := slog.New(
		// зачем ты передаешь env а потом хардкодишь LevelDebug, не порядок
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	return log
}
