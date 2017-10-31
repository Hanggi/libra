package libra

import (
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
	Router  *httprouter.Router
	LRouter LRouter
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

// Static routing
func (app *App) Static(url string, path string) *App {
	Libra.Router.ServeFiles(url+"/*filepath", http.Dir(path))

	return app
}

type middleware struct {
	handler http.Handler
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ctx Context
	ctx.r = r
	ctx.w = w
	for _, value := range Libra.middlewares {
		value(ctx)
	}
	m.handler.ServeHTTP(w, r)
}

// Listen function
func (app *App) Listen(port int) *App {
	if err := http.ListenAndServe(":"+strconv.Itoa(port), &middleware{Libra.Router}); err != nil {
		log.Fatal("Libra listen error: ", err)
	}

	return app
}

// ListenAnd listening functin with and callback
func (app *App) ListenAnd(port int, and func()) *App {
	and()
	app.Listen(port)

	return app
}
