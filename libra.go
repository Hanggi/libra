package libra

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	Context Context
	Log     Log
	// context    controller.LibraContext
	middlewares []Controller
	Session     Session
}

// Exported App struct
var (
	Libra     App
	staticDir map[string]string
)

func init() {
	Libra.Router = httprouter.New()
	Libra.Views = "views"
	Libra.Context.ViewPath = Libra.Views
	Libra.Session.New()
}

// New vv
func New() *App {

	ll := &App{
		Port: 5555,
		//		Config:  Config,
		Router: httprouter.New(),
		LRouter: &LRouter{
			handlers:   nil,
			middleware: middleware{nil, nil},
		},
		Views: "views",
	}

	return ll
}

// Static routing
func (app *App) Static(url string, path string) *App {
	Libra.Router.ServeFiles(url+"/*filepath", http.Dir(path))

	return app
}

/*
 *	middleware module
 */

//type middleware struct {
//	handler http.Handler
//	//	next    *middleware
//}

//// middleware interface imple
//func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	var ctx Context
//	ctx.r = r
//	ctx.w = w
//	for _, value := range Libra.middlewares {
//		value(ctx)
//	}
//	m.handler.ServeHTTP(w, r)
//}

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}

//func build(mws []http.Handler) middleware {
//	var next middleware

//	if len(mws) > 1 {
//		next = build(mws[1:])
//	} else {
//		next = voidMiddleware()
//	}

//	return middleware{mws[0], &next}
//}

//func voidMiddleware() middleware {
//	return middleware{
//		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}),
//		&middleware{},
//	}
//}

//  -----------------------------------------

type Handler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

//type HandlerFunc func(rw http.ResponseWriter, r *http.Request)

//func (h HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
//	h(rw, r)
//}

type HandlerFunc2 func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h HandlerFunc2) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(rw, r, next)
}

type Controller2 func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (c Controller2) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//	var ctx Context
	c(rw, r, next)
	//	Libra.Router.ServeHTTP(rw, r)
}

// Use add middleware
func (l *LRouter) Use(c Controller2) {
	//	var next middleware

	if c == nil {
		fmt.Println("controller can not be nil!")
	}

	l.handlers = append(l.handlers, c)
	l.middleware = build(l.handlers)
}

func build(handlers []Handler) middleware {
	var next middleware

	if len(handlers) == 0 {
		return voidMiddleware()
	} else if len(handlers) > 1 {
		next = build(handlers[1:])
	} else {
		next = voidMiddleware()
	}

	return middleware{handlers[0], &next}
}

func voidMiddleware() middleware {
	return middleware{
		Controller2(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) { Libra.Router.ServeHTTP(rw, r) }),
		&middleware{},
	}
}

type middleware struct {
	handler Handler
	next    *middleware
}

func (m middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if m.handler != nil {
		m.handler.ServeHTTP(rw, r, m.next.ServeHTTP)
	} else {
		fmt.Printf("m.handler nil: %+v \n", m.handler)
	}
	//	Libra.Router.ServeHTTP(rw, r)
}

//  -----------------------------------------
// LRouter vv
type LRouter struct {
	middleware middleware
	handlers   []Handler
}

func (l LRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	l.middleware.ServeHTTP(rw, r)
}

// Listen function
func (app *App) Listen(port int) *App {
	if port <= 0 {
		port = app.Config.Port
	}

	//	http.ListenAndServe()

	server := http.Server{
		Addr:    "127.0.0.1:" + strconv.Itoa(port),
		Handler: app.LRouter,
		//		Handler: &middleware2{Libra.Router},
		//		Handler: exampleMiddleware(Libra.Router),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error eeee")
	}

	return app
}

// ListenAnd listening functin with and callback
func (app *App) ListenAnd(port int, and func()) *App {
	and()
	app.Listen(port)

	return app
}
