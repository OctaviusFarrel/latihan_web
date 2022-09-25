package lib

import (
	"os"

	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Log(fields map[string]interface{}, message string, level int)
}

type Logger struct {
	log *logrus.Logger
}

func NewLogger() ILogger {
	log := logrus.New()

	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{log: log}
}

func (logger *Logger) Log(fields map[string]interface{}, message string, level int) {

	switch level {
	case 0:
		logger.log.WithFields(fields).Info(message)
	case 1:
		logger.log.WithFields(fields).Warning(message)
	case 2:
		logger.log.WithFields(fields).Error(message)
	case 3:
		logger.log.WithFields(fields).Fatal(message)
	}
}
