package logger

import (
	"errors"
	"testing"
)

func TestSlogJsonLogger(t *testing.T) {
	t.Run("SlogJsonLogger tests", TestSlogJsonLogInfo)
}

func TestSlogJsonLogInfo(t *testing.T) {
	t.Run("LogInfo test", func(t *testing.T) {
		InitializeLogger(&SlogJsonLogger{})
		GetLogger().LogInfo("LogInfo test")
	})
}

func TestSlogJsonLogError(t *testing.T) {
	t.Run("LogError test", func(t *testing.T) {
		InitializeLogger(&SlogJsonLogger{})
		GetLogger().LogError("LogError test", errors.New("this is an error"))
	})
}

func TestSlogJsonLogWithFields(t *testing.T) {
	t.Run("LogWithFields", func(t *testing.T) {
		InitializeLogger(&SlogJsonLogger{})
		data := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}
		GetLogger().LogWithFields(data)
	})
}
