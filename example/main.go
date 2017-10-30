package main

import (
	"fmt"

	"github.com/Hanggi/libra/example/route"

	"github.com/Hanggi/libra"
)

var (
	app    = libra.Libra
	router = app.Router
	C      = app.Controller
)

func init() {
	app.Port = 5555
	// route.Init(app)
}

type Person struct {
	UserName string
}

func main() {

	app.Static("/public", "public")

	router.GET("/", C.Route(route.Index))
	router.GET("/test", C.Route(route.Test))
	router.GET("/post", C.Route(route.GetPost))
	router.POST("/post", C.Route(route.Post))

	app.ListenAnd(app.Port, func() {
		fmt.Printf("listening at port: %d", app.Port)
	})
}

// log.WithFields(log.Fields{
// 	"animal": "walrus",
// 	"size":   10,
// }).Info("A group of walrus emerges from the ocean")

// log.WithFields(log.Fields{
// 	"omg":    true,
// 	"number": 122,
// }).Warn("The group's number increased tremendously!")

// log.WithFields(log.Fields{
// 	"omg":    true,
// 	"number": 100,
// }).Fatal("The ice breaks!")

// contextLogger := log.WithFields(log.Fields{
// 	"common": "this is a common field",
// 	"other":  "I also should be logged always",
// })

// contextLogger.Info("I'll be logged with common and other field")
// contextLogger.Info("Me too")
