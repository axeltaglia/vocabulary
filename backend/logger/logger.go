package logger

import "errors"

type Logger interface {
	Init()
	LogError(msg string, err error)
	LogInfo(msg string)
	LogWithFields(fields map[string]interface{})
}

var LogInstance Logger

func InitializeLogger(logger Logger) {
	LogInstance = logger
	LogInstance.Init()
}

func GetLogger() Logger {
	if LogInstance == nil {
		panic(errors.New("logger has not been initialized. Please call InitializeLogger() first"))
	}
	return LogInstance
}
