// Copyright 2017 Hanggi.

package libra

import (
	"github.com/sirupsen/logrus"
)

// Libra service struct
type Libra struct {
	Logrus *logrus.Logger
}

func (l *Libra) Info(args ...interface{}) {
	l.Logrus.Info()
}
