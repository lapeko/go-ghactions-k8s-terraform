package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *logrus.Logger
)

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetLevel(logrus.InfoLevel)
		logger.SetOutput(os.Stdout)
	})

	return logger
}
