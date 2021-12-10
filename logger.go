package libra

import "github.com/sirupsen/logrus"

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}
