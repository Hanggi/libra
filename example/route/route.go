package route

import (
	"fmt"
	"strings"

	"github.com/Hanggi/libra"
	// "github.com/Hanggi/libra/controller"
)

// type API int

var (
	app = libra.Libra
	// context controller.context
)

func init() {

}

// // Page vv
// type Page struct {
// 	Title string
// 	Body  []byte
// }
// func loadPage(title string) (*Page, error) {
// 	filename := title
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }

// Index route
func Index(ctx *libra.Context) {
	fmt.Println("in Index!")
	//	fmt.Printf("Index: %+V\n", ctx)

	ctx.Render("index", nil)

}

type data struct {
	Name string
	Age  int
	Job  string
}

// FormInput ...
func FormInput(ctx *libra.Context) {

	type data struct {
		Title    string
		name     int
		password string
	}

	ctx.Render("form_input", data{"This is form input", 1, "vv"})
}

// FormInputPost ...
func FormInputPost(ctx *libra.Context) {

	//	fmt.Println(ctx.Form)
	type data struct {
		Title    string
		Username string
		Password string
		IsNumber bool
	}

	var isNumber bool
	if ctx.Validate.IsNumber(strings.Join(ctx.Form["password"], "")) {
		isNumber = true
	} else {
		isNumber = false
	}

	ctx.Render("form_input", data{"This is form input", strings.Join(ctx.Form["username"], ""), strings.Join(ctx.Form["password"], ""), isNumber})
}

// Test vv
func Test(ctx *libra.Context) {
	fmt.Println("in Test")

	fmt.Fprintf(ctx.Rw, "Hello")
	//	ctx.Render("index", data{"data's name", 12, "mamada"})
}

// GetPost vv
func GetPost(ctx libra.Context) {
	fmt.Println("in get post")

	ctx.Render("post", data{"data's name", 12, "mamada"})
}

// Post vv
func Post(ctx libra.Context) {
	fmt.Println("in Post")

	// fmt.Println(ctx.Form["username"][0])

	ctx.Render("post", data{ctx.Form["username"][0], 12, "mamada"})
}
