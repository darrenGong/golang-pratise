package main

import (
	"github.com/Sirupsen/logrus"
	"os"
)

func Init() {
	file, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		logrus.WithField("LOGFILE", "log").Fatalf("Failed to open[%s]", "log")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(file)
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	Init()

	logrus.WithFields(logrus.Fields{
		"animals": "walrus",
	}).Info("A walrus appears")
	logrus.WithField("animals", "walrus").Debugf("Debug animals: %s", "walrus")
	logrus.WithField("animals", "walrus").Error("have error")
}
