package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type ILogger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type Logger struct {
	_logger ILogger
}

func (logger *Logger) Init() {
	loggerLevel, err := logrus.ParseLevel(os.Getenv("CORE_CHAINCODE_LOGGING_LEVEL"))
	if err == nil {
		var LoggerLogrus = logrus.New()
		LoggerLogrus.SetLevel(loggerLevel)
		LoggerLogrus.SetFormatter(&logrus.JSONFormatter{})
		LoggerLogrus.SetOutput(os.Stdout)

		logger._logger = LoggerLogrus
	} else {
		logger._logger = &EmptyLogger{}
	}
}

func (logger *Logger) Error(args ...interface{}) {
	logger._logger.Error(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger._logger.Info(args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger._logger.Debug(args...)
}
