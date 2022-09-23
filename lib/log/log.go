package lib

import (
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Log(useCase string, level int)
	String(useCase string, level int) string
}

type Logger struct {
	log *logrus.Logger
}

func NewLogger() ILogger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	return &Logger{log: log}
}

func (logger *Logger) Log(useCase string, level int) {
	field := logrus.Fields{
		"use_case": useCase,
	}

	switch level {
	case 0:
		logger.log.WithFields(field).Info("Test")
	case 1:
		logger.log.WithFields(field).Warning("Test")
	case 2:
		logger.log.WithFields(field).Error("Test")
	case 3:
		logger.log.WithFields(field).Fatal("Test")
	}
}

func (logger *Logger) String(useCase string, level int) (result string) {
	field := logrus.Fields{
		"use_case": useCase,
	}

	result, _ = logger.log.WithFields(field).String()

	return

}
