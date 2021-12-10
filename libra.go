// Copyright 2017 Hanggi.

package libra

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strings"
)

var (
	libra = New()
)

// Libra service struct
type Libra struct {
	Router *gin.Engine
	P      P
}

func (l *Libra) Run() {
	port := viper.GetString("app.port")
	if port != "" {
		_ = l.Router.Run(":" + port)
	} else {
		_ = l.Router.Run()
	}
}

type P struct {
	Logrus *logrus.Logger
	Gin    *gin.Engine
	DB     *gorm.DB
}

func New() *Libra {
	l := &Libra{
		P: P{
			Logrus: logrus.New(),
		},
	}
	l.P.Logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	return l
}

func InitDefaultConfig() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	// Set environment variable prefix.
	viper.SetEnvPrefix("LIBRA")
	if prefix := viper.GetString("app.CONFIG_PREFIX"); prefix != "" {
		viper.SetEnvPrefix(prefix)
	}

	// Environment variable override.
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("default")

	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		Error("Config file not found", err)
	} else if err != nil {
		Error(err)
	}

	// Reads LIBRA_ENV
	env := viper.GetString("env")
	if env != "" {
		// 后合并新的配置文件到默认配置文件
		viper.SetConfigName(env)
		err := viper.MergeInConfig()
		if err != nil {
			Error(err)
			return
		}
	}

	Info("Load config file from " + viper.ConfigFileUsed())
}

func Default() *Libra {
	InitDefaultConfig()

	// Set Gin mode.
	if viper.GetString("app.ginMode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.ForceConsoleColor()
	libra.Router = gin.New()
	libra.Router.Use(gin.Logger(), gin.Recovery())

	if proxies := viper.GetStringSlice("app.trustedProxies"); proxies != nil {
		err := libra.Router.SetTrustedProxies(proxies)
		if err != nil {

		}
	}

	return libra
}
