package libra

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"net/url"
	"path/filepath"
)

type Context struct {
	//xContext

	// w http.ResponseWriter
	// rw         http.ResponseWriter
	// r          *http.Request
	Rw         http.ResponseWriter
	R          *http.Request
	Ps         httprouter.Params
	ViewPath   string
	Method     string
	URL        *url.URL
	Proto      string
	ProtoMajor string
	IPp        string
	Form       url.Values
	Validate   FormUtil
	Query      url.Values
	GetParam   func(string) string

	app         *Libra
	index       int8
	controllers []Controller
	router      *httprouter.Router
}

func (ctx *Context) reset() {
	ctx.index = -1
	ctx.controllers = nil
}

func (ctx *Context) Next() {
	ctx.index++
	if ctx.index < int8(len(ctx.controllers)) {
		ctx.controllers[ctx.index](ctx)
	} else {
		ctx.router.ServeHTTP(ctx.Rw, ctx.R)
	}
}

func (ctx *Context) Render(view string, data interface{}) {
	// fmt.Println("in Render")
	path := filepath.Join(ctx.app.Views, view+".html")
	fmt.Println(path)

	t, err := template.ParseFiles(path)

	if err != nil {
		fmt.Println("【Error】", err)
		ctx.Rw.WriteHeader(http.StatusInternalServerError) // Proper HTTP response
		return
	}
	t.Execute(ctx.Rw, data)
}

func (ctx *Context) GetHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx.Rw = w
	ctx.R = r
	ctx.Ps = ps
}

func (ctx *Context) Param(key string) string {
	return ctx.Ps.ByName(key)
}

// Query: shortcut
func (ctx *Context) Query(name string) string {
	return ctx.R.URL.Query().Get(name)
}
