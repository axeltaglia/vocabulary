package slogJsonLogger

import (
	"errors"
	"testing"
	"vocabulary/logger"
)

func TestSlogJsonLogger(t *testing.T) {
	t.Run("SlogJsonLogger tests", TestLogInfo)
}

func TestLogInfo(t *testing.T) {
	t.Run("LogInfo test", func(t *testing.T) {
		logger.InitializeLogger(&SlogJsonLogger{})
		logger.GetLogger().LogInfo("LogInfo test")
	})
}

func TestLogError(t *testing.T) {
	t.Run("LogError test", func(t *testing.T) {
		logger.InitializeLogger(&SlogJsonLogger{})
		logger.GetLogger().LogError("LogError test", errors.New("this is an error"))
	})
}

func TestLogWithFields(t *testing.T) {
	t.Run("LogWithFields", func(t *testing.T) {
		logger.InitializeLogger(&SlogJsonLogger{})
		data := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}
		logger.GetLogger().LogWithFields(data)
	})
}
