package log

//More log example
//https://github.com/Sirupsen/logrus#example

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func LogError(r *http.Request, err error, info string, logger *log.Logger) {
	logger.WithFields(log.Fields{
		"error":  err.Error(),
		"method": r.Method,
		"url":    r.URL.String(),
	}).Error(info)
}

// logger.WithFields(log.MySQLError(err.Error(), eatery.ToString())).Error("Error in Creating Eatery")
func Error(msg string) map[string]interface{} {
	return log.Fields{"error": msg}
}

func Debug(msg string) map[string]interface{} {
	return log.Fields{"debug": msg}
}

func Info(msg string) map[string]interface{} {
	return log.Fields{"info": msg}
}

func Warn(msg string) map[string]interface{} {
	return log.Fields{"warn": msg}
}

//Warning: this will call os.Exit(1) after logged
func Fatal(msg string) map[string]interface{} {
	return log.Fields{"fatal": msg}
}

// Calls panic() after logged
func Panic(msg string) map[string]interface{} {
	return log.Fields{"panic": msg}
}
