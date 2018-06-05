// Copyright 2017 Hanggi.

package libra

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"fmt"
	"github.com/julienschmidt/httprouter"
)

// App struct
type Libra struct {
	LRouter

	Port    int
	Config  Config
	Router  *httprouter.Router
	Views   string
	Context *Context
	Log     Log
	Session Session
	pool    sync.Pool
	//LRouter *LRouter
	// context    controller.LibraContext
	//	middlewares []Controller
	// Controller Controller
}

// Exported App struct
var (
	//Libra App
	DEBUG bool = true
	Debug      = LibraDebug{}

//	staticDir map[string]string
)

func init() {

}

// New function return a app instance with context allocated
func New() *Libra {

	ll := &Libra{
		Port: 5555,
		//		Config:  Config,
		//Router: httprouter.New(),
		LRouter: LRouter{
			nil,
			nil,
			httprouter.New(),
		},
		Views: "views",
	}
	ll.LRouter.app = ll
	ll.pool.New = func() interface{} {
		return ll.allocateContext()
	}

	return ll
}

// allocateContext function return a context instance into sync.pool
func (app *Libra) allocateContext() *Context {
	return &Context{app: app}
}

// Listen function
func (app *Libra) Listen(port int) *Libra {
	if port <= 0 {
		port = app.Config.Port
	}

	server := http.Server{
		Addr:    "127.0.0.1:" + strconv.Itoa(port),
		Handler: app, // app implemented ServeHTTP
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Listen error: ", err)
	}

	return app
}

// ListenAnd listening functin with and callback
func (app *Libra) ListenAnd(port int, callback func()) *Libra {
	callback()
	app.Listen(port)

	return app
}

func (app *Libra) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Debug.Print("Serve HTTP of app")
	fmt.Println("Serve HTTP of app")
	//println("Serve HTTP of app")
	ctx := app.pool.Get().(*Context)
	//	V("ServeHTTP ctx: ", ctx)
	// ctx.rw = w
	ctx.Rw = w
	ctx.R = r

	ctx.Method = r.Method
	ctx.URL = r.URL

	ctx.reset()
	ctx.router = app.LRouter.router
	// Here, needs init

	app.handleHTTPRequest(ctx)

	app.pool.Put(ctx)
}
