// Copyright 2017 Hanggi.

package libra

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	libra = New()
)

// Libra service struct
type Libra struct {
	Logrus *logrus.Logger
	Gin    *gin.Engine
	DB     *gorm.DB
}

func New() *Libra {
	l := &Libra{
		Logrus: logrus.New(),
	}
	l.Logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	return l
}
