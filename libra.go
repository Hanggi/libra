// Copyright 2017 Hanggi.

package libra

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	// "./context"
	// "github.com/Hanggi/libra/controller"

	"github.com/julienschmidt/httprouter"
)

// App struct
type App struct {
	Port    int
	Config  Config
	Router  *httprouter.Router
	LRouter *LRouter
	Views   string
	// Controller Controller
	Context *Context
	Log     Log
	// context    controller.LibraContext
	//	middlewares []Controller
	Session Session
	pool    sync.Pool
}

// Exported App struct
var (
	Libra App
	DEBUG bool = true

//	staticDir map[string]string
)

func init() {
	//	Libra.Router = httprouter.New()
	//	Libra.Views = "views"
	//	Libra.Context.ViewPath = Libra.Views
	//	Libra.Session.New()
}

// New function return a app instance with context allocated
func New() *App {

	ll := &App{
		Port: 5555,
		//		Config:  Config,
		Router: httprouter.New(),
		LRouter: &LRouter{
			handlers: nil,
			//			middleware: middleware{nil, nil},
		},
		Views: "views",
	}
	ll.pool.New = func() interface{} {
		return ll.allocateContext()
	}

	return ll
}

// allocateContext function return a context instance into sync.pool
func (app *App) allocateContext() *Context {
	return &Context{app: app}
}

// Static function set the static file path with path string
func (app *App) Static(url string, path string) *App {
	app.Router.ServeFiles(url+"/*filepath", http.Dir(path))

	return app
}

func (app *App) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//	DebugP("App ServeHTTP")

	ctx := app.pool.Get().(*Context)
	//	V("ServeHTTP ctx: ", ctx)
	ctx.rw = rw
	ctx.w = rw
	ctx.r = r

	ctx.Method = r.Method
	ctx.URL = r.URL

	ctx.reset()
	ctx.router = app.Router
	// Here, needs init

	app.handleHTTPRequest(ctx)

	app.pool.Put(ctx)

}

func (app *App) handleHTTPRequest(ctx *Context) {
	//	DebugP("handleHTTPRequest")
	//	httpMethod := ctx.r.Method
	//	path := ctx.r.URL
	app.LRouter.handleHTTPRequest(ctx)
}

func (app *App) Use(middles ...Controller) {
	app.LRouter.Use(middles...)
}

func (app *App) GET(path string, c Controller) {
	app.Router.GET(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := app.pool.Get().(*Context)
		ctx.w = w
		ctx.rw = w
		ctx.r = r
		ctx.Rw = w
		ctx.R = r

		c(ctx)

		app.pool.Put(ctx)
	})
}

// Listen function
func (app *App) Listen(port int) *App {
	if port <= 0 {
		port = app.Config.Port
	}

	server := http.Server{
		Addr:    "127.0.0.1:" + strconv.Itoa(port),
		Handler: app,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Listen error: ", err)
	}

	return app
}

// ListenAnd listening functin with and callback
func (app *App) ListenAnd(port int, and func()) *App {
	and()
	app.Listen(port)

	return app
}
