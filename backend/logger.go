// logger.go
package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	customFormatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2000-01-02 15:04:05", // format
	}

	Logger.SetFormatter(customFormatter)
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
}
