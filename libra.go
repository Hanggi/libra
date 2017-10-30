package libra

import (
	"log"
	"net/http"
	"strconv"

	// "./context"
	"github.com/Hanggi/libra/controller"
	"github.com/julienschmidt/httprouter"
)

// App struct
type App struct {
	Port       int
	Router     *httprouter.Router
	Views      string
	Controller controller.Controller
	// context    controller.LibraContext
}

// Exported App struct
var (
	Libra     App
	staticDir map[string]string
)

func init() {
	Libra.Router = httprouter.New()
	Libra.Views = "views"
	controller.LibraContext.ViewPath = Libra.Views
}

type Person struct {
	UserName string
}

// Static routing
func (app *App) Static(url string, path string) *App {
	Libra.Router.ServeFiles(url+"/*filepath", http.Dir(path))

	return app
}

// Listen function
func (app *App) Listen(port int) *App {
	if err := http.ListenAndServe(":"+strconv.Itoa(port), Libra.Router); err != nil {
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

// func Test() {
// 	fmt.Println("lib test")
// }

// json := `{"data": {"a": "json str2", "b": 2, "c": [1, 2, 3]}}`
// js, err := simplejson.NewJson([]byte(json))
// if err != nil {
// 	fmt.Println("err json")
// }
// a := js.Get("data").Get("a").MustString()
// fmt.Println(a)

// t.Execute(w, Person{UserName: a})
