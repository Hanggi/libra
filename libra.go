// Copyright 2017 Hanggi.

package libra

import (
	"github.com/sirupsen/logrus"
)

var (
	libra = New()
)

// Libra service struct
type Libra struct {
	Logrus *logrus.Logger
}

func New() *Libra {
	return &Libra{}
}

func Info(args ...interface{}) {
	libra.Logrus.Info()
}
