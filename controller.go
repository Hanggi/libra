package libra

//
//import (
//	"fmt"
//	"html/template"
//	"net/http"
//	"net/url"
//	"time"
//
//	// "../../libra"
//
//	"github.com/julienschmidt/httprouter"
//	"golang.org/x/net/context"
//)
//
//// Context : http route context
//type xContext = context.Context
//
//func init() {
//
//	//	fmt.Printf("%+v \n", CC.)
//
//}
//
//// Context ...
//type Context struct {
//	xContext
//
//	// w http.ResponseWriter
//	// rw         http.ResponseWriter
//	// r          *http.Request
//	Rw         http.ResponseWriter
//	R          *http.Request
//	ps         httprouter.Params
//	ViewPath   string
//	Method     string
//	URL        *url.URL
//	Proto      string
//	ProtoMajor string
//	IPp        string
//	Form       url.Values
//	Validate   FormUtil
//	Query      url.Values
//	GetParam   func(string) string
//
//	app      *Libra
//	index    int8
//	handlers []Controller
//	router   *httprouter.Router
//}
//
//func (ctx *Context) reset() {
//	ctx.index = -1
//	ctx.handlers = nil
//}
//
//// Render in context
//func (ctx *Context) Render(view string, data interface{}) {
//	// fmt.Println("in Render")
//	path := ctx.app.Views + "/" + view + ".html"
//
//	t, err := template.ParseFiles(path)
//
//	if err != nil {
//		fmt.Println("【Error】", err)
//		ctx.Rw.WriteHeader(http.StatusInternalServerError) // Proper HTTP response
//		return
//	}
//	t.Execute(ctx.Rw, data)
//}
//
//// SetCookie vv
//func (ctx *Context) SetCookie(name string, value string, day int) {
//	expiration := time.Now()
//	expiration = expiration.AddDate(0, 0, day)
//	cookie := http.Cookie{Name: name, Value: value, Expires: expiration}
//	http.SetCookie(ctx.Rw, &cookie)
//}
//
//// GetCookie get single cookie by name
//func (ctx *Context) GetCookie(name string) *http.Cookie {
//	cookie, err := ctx.R.Cookie(name)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	// fmt.Fprint(ctx.w, cookie.Name)
//	return cookie
//}
//
//// Controller vv
//type Controller func(ctx *Context)
//
//// Next ...
//func (ctx *Context) Next() {
//	ctx.index++
//	//	V("ctx.handlers: ", ctx.handlers)
//	if ctx.index < int8(len(ctx.handlers)) {
//		ctx.handlers[ctx.index](ctx)
//	} else {
//		//		httprouter.serveHTTP()
//		ctx.router.ServeHTTP(ctx.Rw, ctx.R)
//	}
//}
//
///*
// * LRouter which support middleware with next function without any warpper
// */
//
//// LRouter ...
//type LRouter struct {
//	//	middleware middleware
//	handlers []Controller
//	app      *Libra
//}
//
//func (l *LRouter) handleHTTPRequest(ctx *Context) {
//	println("lrouter http request")
//	ctx.index = 0
//
//	if l.handlers != nil {
//		//		l.handlers
//		ctx.handlers = l.handlers
//		l.handlers[0](ctx)
//	} else {
//		ctx.router.ServeHTTP(ctx.Rw, ctx.R)
//	}
//}
//
////// middlware struct
////type middleware struct {
////	//	handler ContextHandler
////	next *middleware
////}
//
//// Use add middleware
//func (l *LRouter) Use(c ...Controller) {
//	//	var next middleware
//
//	if c == nil {
//		fmt.Println("controller can not be nil!")
//	}
//
//	l.handlers = append(l.handlers, c...)
//	//	V("handlers", l.handlers)
//	//	l.middleware = build(l.handlers)
//}
//
//func setContext(ctx *Context, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	// ctx.w = w
//	ctx.R = r
//	// ctx.ps = ps
//	ctx.Rw = w
//
//	ctx.Method = r.Method
//	ctx.URL = r.URL
//	ctx.IPp = r.RemoteAddr
//	ctx.Query = r.URL.Query()
//	ctx.GetParam = ps.ByName
//}
//
//// GET http get request
////func (app *Libra) GET(path string, c Controller) *Libra {
////	// app.Router.GET(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
////	// 	ctx := app.pool.Get().(*Context)
////	// 	ctx.w = w
////	// 	ctx.rw = w
////	// 	ctx.r = r
////	// 	ctx.Rw = w
////	// 	ctx.R = r
////
////	// 	c(ctx)
////
////	// 	app.pool.Put(ctx)
////	// })
////	app.LRouter.GET(path, c)
////	return app
////}
//
//// GET vv
//func (l *LRouter) GET(path string, controller Controller) {
//	ctx := l.app.pool.Get().(*Context)
//
//	l.app.Router.GET(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//
//		println("!!!!!!!!!!!!!")
//		setContext(ctx, w, r, ps)
//		// fmt.Printf("%+v \n", r)
//		// fmt.Println(ctx.Query)
//		// fmt.Println(ps)
//		// fmt.Println(ctx.GetParam("id"))
//		controller(ctx)
//	})
//}
//
//// POST vv
//func (l *LRouter) POST(path string, controller Controller) {
//	var ctx *Context
//
//	l.app.Router.POST(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//		// defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
//		// 	log.WithFields(log.Fields{
//		// 		// "Method": ctx.Method + " " + ctx.URL.Path,
//		// 		"time": d,
//		// 	}).Info(ctx.Method + " " + ctx.URL.Path)
//		// })
//
//		setContext(ctx, w, r, ps)
//		if ctx.Method == "POST" {
//			r.ParseForm()
//			ctx.Form = r.Form
//			//			fmt.Println(ctx.Form)
//		}
//
//		controller(ctx)
//	})
//}
//
//// PUT vv
////func (l *LRouter) PUT(path string, controller Controller) {
////	var ctx *Context
////
////	Libra.Router.GET(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
////		defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
////			log.WithFields(log.Fields{
////				// "Method": ctx.Method + " " + ctx.URL.Path,
////				"time": d,
////			}).Info(ctx.Method + " " + ctx.URL.Path)
////		})
////
////		setContext(ctx, w, r, ps)
////
////		controller(ctx)
////	})
////}
////
////// DELETE vv
////func (l *LRouter) DELETE(path string, controller Controller) {
////	var ctx *Context
////
////	Libra.Router.DELETE(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
////		defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
////			log.WithFields(log.Fields{
////				// "Method": ctx.Method + " " + ctx.URL.Path,
////				"time": d,
////			}).Info(ctx.Method + " " + ctx.URL.Path)
////		})
////
////		setContext(ctx, w, r, ps)
////
////		controller(ctx)
////	})
////}
