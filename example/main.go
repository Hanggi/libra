package main

import (
	"fmt"

	"github.com/Hanggi/libra"
	"github.com/Hanggi/libra/example/route"
)

var (
	app    = libra.Libra
	router = app.LRouter
)

func init() {
	app.Port = 5555
}

type data struct {
	Name string
	Age  int
	Job  string
}

func middleware() {
	fmt.Println("fff")
}

func main() {

	app.Static("/public", "public")

	router.Use(middleware)

	router.GET("/", route.Index)
	router.GET("/test", route.Test)
	router.GET("/test/:id", route.Test)
	router.GET("/test/:id/*action", route.Test)
	router.GET("/post", func(ctx libra.Context) {
		fmt.Println("in get post")

		ctx.Render("post", data{"data's name", 12, "mamada"})
	})

	app.ListenAnd(app.Port, func() {
		fmt.Printf("listening at port: %d\n", app.Port)
	})
}
