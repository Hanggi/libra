package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	// "../../libra"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// type Controller struct {
// 	// Ctx       *Context
// 	// Tpl       *template.Template
// 	Data      map[interface{}]interface{}
// 	ChildName string
// 	TplNames  string
// 	Layout    []string
// 	TplExt    string
// }

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

		// fmt.Println("\n", r, "[\n]")

		fmt.Printf("%+v", r)
		fmt.Println(ps, "\n!!")

		// PrettyPrint(r)

		log.WithFields(log.Fields{
			"Method": ctx.Method,
			"size":   10,
		}).Info("log test")

		route(ctx)
	}
}

// func PrettyPrint(v interface{}) {
// 	b, _ := json.MarshalIndent(v, "", "  ")
// 	fmt.Println(string(b))
// }
