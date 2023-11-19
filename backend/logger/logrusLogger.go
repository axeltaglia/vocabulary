package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func (o *LogrusLogger) Init() {
	o.log = logrus.New()

	o.log.Out = os.Stdout
	//log.SetFormatter(&logrus.JSONFormatter{})
	o.log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	//file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//	log.Out = file
	//} else {
	//	log.Info("Failed to log to file, using default stderr")
	//}

}

func (o *LogrusLogger) LogError(msg string, err error) {
	_, file, line, _ := runtime.Caller(1)
	fileParts := strings.Split(file, "/")
	fileName := fileParts[len(fileParts)-1]

	o.log.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     "error",
		"message":   msg,
		"error":     err.Error(),
		"file":      fileName,
		"line":      line,
	}).Error("Error occurred")
	o.log.Error(getStackTrace(err))
}

func (o *LogrusLogger) LogInfo(msg string) {
	_, file, line, _ := runtime.Caller(1)
	fileParts := strings.Split(file, "/")
	fileName := fileParts[len(fileParts)-1]

	o.log.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     "info",
		"message":   msg,
		"file":      fileName,
		"line":      line,
	}).Info("Info message")

}

func (o *LogrusLogger) LogWithFields(fields map[string]interface{}) {
	o.log.WithFields(fields).Error("Error occurred")
}

func getStackTrace(err error) string {
	stackTrace := string(debug.Stack())
	return fmt.Sprintf("%v\n%s\n", err.Error(), stackTrace)
}
