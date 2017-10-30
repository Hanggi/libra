package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	// "../../libra"
	"github.com/Hanggi/libra/util"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// Context ..
/*Context struct
 *	http context
 */
type Context struct {
	w          http.ResponseWriter
	r          *http.Request
	ps         httprouter.Params
	ViewPath   string
	Method     string
	URL        *url.URL
	Proto      string
	ProtoMajor string
	IPp        string
	Form       url.Values
}

var (
	LibraContext Context
)

// Render in context
func (ctx *Context) Render(view string, data interface{}) {
	fmt.Println("in Render")

	// fmt.Println(LibraContext.ViewPath)
	t, err := template.ParseFiles(LibraContext.ViewPath + "/" + view + ".html")
	// t, err := template.ParseFiles(view + ".html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(ctx.w, data)
}

/*Controller struct
 *	A wrapper of httprouter Handle
 */
type Controller struct {
}

// Route controller wrap
func (c *Controller) Route(route func(ctx_para Context)) httprouter.Handle {
	var ctx Context

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx.w = w
		ctx.r = r
		ctx.ps = ps

		ctx.Method = r.Method
		ctx.URL = r.URL
		ctx.IPp = r.RemoteAddr

		fmt.Printf("%+v \n\n", r)
		// fmt.Println(ps, "\n!!")

		if ctx.Method == "POST" {
			r.ParseForm()
			ctx.Form = r.Form
		}

		start := time.Now()
		defer util.CalcTimeEnd(start, func(d time.Duration) {
			log.WithFields(log.Fields{
				"Method": ctx.Method,
				"time":   d,
			}).Info("log test")
		})

		route(ctx)

	}
}
