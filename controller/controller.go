package controller

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

// Context ..
/*Context struct
 *	http context
 */

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
}

// var
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
		// return false
	}
	// fmt.Fprint(ctx.w, cookie.Name)
	return cookie
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

		// fmt.Println(ctx.Validate.IsEmail("aef@f2.com"))
		// fmt.Println(ctx.Validate.IsEmail("aef@f@2.com"))

		ctx.SetCookie("thisisname", "this is value", 1)
		fmt.Println(ctx.GetCookie("thisisname"))

		if ctx.Method == "POST" {
			r.ParseForm()
			ctx.Form = r.Form
			fmt.Printf("%+v", ctx.Form)
		}

		defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
			log.WithFields(log.Fields{
				"Method": ctx.Method + " " + ctx.URL.Path,
				"time":   d,
			}).Info("log test")
		})

		route(ctx)

	}
}
