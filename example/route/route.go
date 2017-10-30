package route

import (
	"fmt"

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
func Index(ctx libra.Context) {
	fmt.Println("in Index")

	ctx.Render("index", nil)

}

type data struct {
	Name string
	Age  int
	Job  string
}

// Test vv
func Test(ctx libra.Context) {
	fmt.Println("in Test")

	ctx.Render("index", data{"data's name", 12, "mamada"})
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
