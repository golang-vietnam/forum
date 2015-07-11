package log

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() *logrus.Logger {
	return &logrus.Logger{
		Out: &lumberjack.Logger{
			Filename:   "log/foo.log",
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		},
		Formatter: new(logrus.JSONFormatter),
		Level:     logrus.DebugLevel,
	}
}
