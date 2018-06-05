package libra

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IRouter interface {
	IRoutes
}
type IRoutes interface {
	Use(...Controller) IRoutes

	GET(string, Controller) IRoutes
	POST(string, Controller) IRoutes
}

type Controller func(ctx *Context)

type LRouter struct {
	controllers []Controller
	app         *Libra
	router      *httprouter.Router
}

func init() {

}

func (l *LRouter) handleHTTPRequest(ctx *Context) {
	//Debug.Print("lrouter http request")
	fmt.Println("lrouter http request")
	ctx.index = 0

	if l.controllers != nil {
		//		l.handlers
		ctx.controllers = l.controllers
		l.controllers[0](ctx)
	} else {
		ctx.router.ServeHTTP(ctx.Rw, ctx.R)
	}
}

// GET http get request
func (l *LRouter) GET(path string, c Controller) IRoutes {
	l.router.GET(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := l.app.pool.Get().(*Context)
		ctx.GetHandle(w, r, ps)

		c(ctx)

		l.app.pool.Put(ctx)
	})
	return l
}

func (l *LRouter) POST(path string, c Controller) IRoutes {
	l.router.POST(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := l.app.pool.Get().(*Context)
		ctx.GetHandle(w, r, ps)

		c(ctx)

		l.app.pool.Put(ctx)
	})
	return l
}

func (l *LRouter) handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) IRoutes {

	return l
}

func (l *LRouter) Static(url string, path string) IRoutes {
	l.router.ServeFiles(url+"/*filepath", http.Dir(path))

	return l
}

// middlware struct
type middleware struct {
	//	handler ContextHandler
	next *middleware
}

func (l *LRouter) Use(c ...Controller) IRoutes {
	//	var next middleware

	if c == nil {
		fmt.Println("controller can not be nil!")
	}

	l.controllers = append(l.controllers, c...)
	//	V("handlers", l.handlers)
	//	l.middleware = build(l.handlers)

	return l
}
