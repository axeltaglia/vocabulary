package logger

import (
	"log/slog"
	"os"
)

type SlogJsonLogger struct {
	logger *slog.Logger
}

func (o *SlogJsonLogger) Init() {
	o.logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(o.logger)
}

func (o *SlogJsonLogger) LogError(msg string, err error) {
	o.logger.Error(msg, err)
}

func (o *SlogJsonLogger) LogInfo(msg string) {
	o.logger.Info(msg)
}

func (o *SlogJsonLogger) LogWarn(msg string) {
	o.logger.Warn(msg)
}

func (o *SlogJsonLogger) LogWithFields(fields map[string]interface{}) {
	o.logger.Info("msg", "data", fields)
}
