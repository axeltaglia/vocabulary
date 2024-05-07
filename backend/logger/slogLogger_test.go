package logger

import (
	"errors"
	"testing"
)

func TestSlogLogger(t *testing.T) {
	t.Run("SlogLogger tests", TestLogInfo)
}

func TestLogInfo(t *testing.T) {
	t.Run("LogInfo test", func(t *testing.T) {
		InitializeLogger(&SlogLogger{})
		GetLogger().LogInfo("LogInfo test")
	})
}

func TestLogError(t *testing.T) {
	t.Run("LogError test", func(t *testing.T) {
		InitializeLogger(&SlogLogger{})
		GetLogger().LogError("LogError test", errors.New("this is an error"))
	})
}

func TestLogWithFields(t *testing.T) {
	t.Run("LogWithFields", func(t *testing.T) {
		InitializeLogger(&SlogLogger{})
		data := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}
		GetLogger().LogWithFields(data)
	})
}
