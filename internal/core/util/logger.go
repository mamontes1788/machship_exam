package util

import (
	"github.com/sirupsen/logrus"
)

func logger() *logrus.Logger {
	logger := logrus.New()
	return logger
}

func Errorf(format string, args ...interface{}) {
	logger().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger().Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger().Panicf(format, args...)
}

func Errorln(args ...interface{}) {
	logger().Errorln(args...)
}

func Fatalln(format string, args ...interface{}) {
	logger().Fatalln(args...)
}

func Panicln(format string, args ...interface{}) {
	logger().Panicln(args...)
}

func Infof(format string, args ...interface{}) {
	logger().Infof(format, args...)
}

func Infoln(args ...interface{}) {
	logger().Infoln(args...)
}
