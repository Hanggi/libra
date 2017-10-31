package libra

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	// "../../libra"

	"github.com/Hanggi/libra/util"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type formUtil struct {
}

func (f formUtil) IsNumber(str string) bool {
	getint, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if getint > 100 {
		//太大了
	}

	return true
}

func (f formUtil) IsHan(str string) bool {
	if m, err := regexp.MatchString("^\\p{Han}+$", str); !m {
		if err != nil {
			fmt.Println(err)
			return false
		}
		return false
	}

	return true
}

func (f formUtil) IsEng(str string) bool {
	if m, err := regexp.MatchString("^[a-zA-Z]+$", str); !m {
		if err != nil {
			fmt.Println(err)
			return false
		}
		return false
	}

	return true
}

func (f formUtil) IsEmail(str string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, str); !m {
		return false
	}

	return true
}

// Context : http route context
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
	Validate   formUtil
	Query      url.Values
	GetParam   func(string) string
}

// Render in context
func (ctx *Context) Render(view string, data interface{}) {
	// fmt.Println("in Render")

	t, err := template.ParseFiles(Libra.Context.ViewPath + "/" + view + ".html")

	if err != nil {
		fmt.Println("【Error】", err)
		ctx.w.WriteHeader(http.StatusInternalServerError) // Proper HTTP response
		return
	}
	t.Execute(ctx.w, data)
}

// SetCookie vv
func (ctx *Context) SetCookie(name string, value string, day int) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, day)
	cookie := http.Cookie{Name: name, Value: value, Expires: expiration}
	http.SetCookie(ctx.w, &cookie)
}

// GetCookie get single cookie by name
func (ctx *Context) GetCookie(name string) *http.Cookie {
	cookie, err := ctx.r.Cookie(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// fmt.Fprint(ctx.w, cookie.Name)
	return cookie
}

// LRouter vv
type LRouter struct {
}

func setContext(ctx *Context, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx.w = w
	ctx.r = r
	ctx.ps = ps

	ctx.Method = r.Method
	ctx.URL = r.URL
	ctx.IPp = r.RemoteAddr
	ctx.Query = r.URL.Query()
	ctx.GetParam = ps.ByName
}

type Controller func()

func (r *LRouter) Use(c Controller) {
	Libra.middlewares = append(Libra.middlewares, c)
}

// GET vv
func (r *LRouter) GET(path string, controller func(ctx_para Context)) {
	var ctx Context

	Libra.Router.GET(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
			log.WithFields(log.Fields{
				// "Method": ctx.Method + " " + ctx.URL.Path,
				"time": d,
			}).Info("[" + ctx.Method + "] " + ctx.URL.Path)
		})

		setContext(&ctx, w, r, ps)
		// fmt.Printf("%+v \n", r)
		// fmt.Println(ctx.Query)
		// fmt.Println(ps)
		// fmt.Println(ctx.GetParam("id"))

		controller(ctx)
	})
}

// POST vv
func (r *LRouter) POST(path string, controller func(ctx_para Context)) {
	var ctx Context

	Libra.Router.POST(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
			log.WithFields(log.Fields{
				// "Method": ctx.Method + " " + ctx.URL.Path,
				"time": d,
			}).Info(ctx.Method + " " + ctx.URL.Path)
		})

		if ctx.Method == "POST" {
			r.ParseForm()
			ctx.Form = r.Form
			fmt.Printf("%+v", ctx.Form)
		}

		controller(ctx)
	})
}

// PUT vv
func (r *LRouter) PUT(path string, controller func(ctx_para Context)) {
	var ctx Context

	Libra.Router.GET(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
			log.WithFields(log.Fields{
				// "Method": ctx.Method + " " + ctx.URL.Path,
				"time": d,
			}).Info(ctx.Method + " " + ctx.URL.Path)
		})

		setContext(&ctx, w, r, ps)

		controller(ctx)
	})
}

// DELETE vv
func (r *LRouter) DELETE(path string, controller func(ctx_para Context)) {
	var ctx Context

	Libra.Router.DELETE(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
			log.WithFields(log.Fields{
				// "Method": ctx.Method + " " + ctx.URL.Path,
				"time": d,
			}).Info(ctx.Method + " " + ctx.URL.Path)
		})

		setContext(&ctx, w, r, ps)

		controller(ctx)
	})
}
