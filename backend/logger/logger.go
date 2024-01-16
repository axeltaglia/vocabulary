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

var log = logrus.New()

func InitLogger() {
	log.Out = os.Stdout
	//log.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(&logrus.TextFormatter{
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

func Log() *logrus.Logger {
	return log
}

// LogError logs an error message along with structured information
func LogError(msg string, err error) {
	_, file, line, _ := runtime.Caller(1)
	fileParts := strings.Split(file, "/")
	fileName := fileParts[len(fileParts)-1]

	log.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     "error",
		"message":   msg,
		"error":     err.Error(),
		"file":      fileName,
		"line":      line,
	}).Error("Error occurred")
	log.Error(getStackTrace(err))
}

// LogInfo logs an info message along with structured information
func LogInfo(msg string) {
	_, file, line, _ := runtime.Caller(1)
	fileParts := strings.Split(file, "/")
	fileName := fileParts[len(fileParts)-1]

	log.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     "info",
		"message":   msg,
		"file":      fileName,
		"line":      line,
	}).Info("Info message")
}

func getStackTrace1() string {
	var stackTrace strings.Builder
	stackTrace.WriteString("\nStack Trace:\n")

	for i := 3; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		stackTrace.WriteString(fmt.Sprintf("%s:%d\n", file, line))
	}

	return stackTrace.String()
}

func getStackTrace(err error) string {
	stackTrace := string(debug.Stack())
	return fmt.Sprintf("%v\n%s\n", err.Error(), stackTrace)
}
