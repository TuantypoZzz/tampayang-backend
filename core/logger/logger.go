package mylogger

import (
	"os"

	"tampayang-backend/config"

	"github.com/sirupsen/logrus"
)

func Info(action string, data interface{}) {
	logger := logrus.New()
	path := logPath() + "/application.log"
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"action": action,
		"data":   data,
	}).Info("Logger Info")
}

func Error(action string, data interface{}) {
	logger := logrus.New()
	path := logPath() + "/application.log"
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"action": action,
		"data":   data,
	}).Error("Logger Error")
}

func Trace(action string, data interface{}) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	path := logPath() + "/application.log"
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"action": action,
		"data":   data,
	}).Trace("Logger Trace")
}

func logPath() string {
	rootPath := config.ProjectRootPath
	logPath := rootPath + "/logs"

	return logPath
}
