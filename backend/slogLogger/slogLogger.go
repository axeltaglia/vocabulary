package slogLogger

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
)

type SlogLogger struct {
	logger *slog.Logger
}

func (o *SlogLogger) Init() {
	o.logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(o.logger)
}

func (o *SlogLogger) LogError(msg string, err error) {
	o.logger.Error(msg, err)
}

func (o *SlogLogger) LogInfo(msg string) {
	o.logger.Info(msg)
}

func (o *SlogLogger) LogWarn(msg string) {
	o.logger.Warn(msg)
}

func (o *SlogLogger) LogWithFields(fields map[string]interface{}) {
	o.logger.Info("msg", "data", fields)
}

func getStackTrace(err error) string {
	stackTrace := string(debug.Stack())
	return fmt.Sprintf("%v\n%s\n", err.Error(), stackTrace)
}
