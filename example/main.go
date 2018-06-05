package main

import (
	"fmt"

	"github.com/Hanggi/libra"
	// _ "github.com/go-sql-driver/mysql"
)

var (
// app    = libra.Libra
// router = app.LRouter
)

func init() {
	// app.Port = 5555
}

type data struct {
	Name string
	Age  int
	Job  string
}

func middleware(ctx *libra.Context) {
	fmt.Println("This is a 111 middleware!")

	ctx.Next()
	fmt.Println("middleware 111 end!!!")
}
func middleware2(ctx *libra.Context) {
	fmt.Println("This is a middleware22222!")

	ctx.Next()
	fmt.Println("This is a middleware22222  end!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := libra.New()
	//	port := app.Config.Port
	port := 5555

	//	router := app.LRouter

	app.Static("/public", "public")
	app.Use(libra.Logger())

	// db, err := sql.Open("mysql", "root:qq110119120@/fortest")
	// checkErr(err)

	// stmt, err := db.Prepare("INSERT test1 SET username=?,age=?,salary=?")
	// checkErr(err)

	// res, err := stmt.Exec("ff", "235", "1129")1
	// checkErr(err)

	// id, err := res.LastInsertId()

	// fmt.Println(id)

	app.Use(middleware)
	app.Use(middleware2)

	// app.GET("/form/input", route.FormInput)
	// app.POST("/form/input", route.FormInputPost)

	app.GET("/", Index)
	app.GET("/test", Test)

	// Parameters in path
	app.GET("/param/:id", Param)

	app.ListenAnd(port, func() {
		fmt.Printf("listening at port: %d\n", port)
	})
}

// Index route
func Index(ctx *libra.Context) {
	fmt.Println("in Index!")
	//	fmt.Printf("Index: %+V\n", ctx)

	ctx.Render("index", nil)

}

// Test vv
func Test(ctx *libra.Context) {
	fmt.Println("in Test")

	fmt.Fprintf(ctx.Rw, "Hello")
	//	ctx.Render("index", data{"data's name", 12, "mamada"})
}

func Param(ctx *libra.Context) {
	fmt.Println("in Path param")
	fmt.Println(ctx.Ps)

	fmt.Fprintf(ctx.Rw, "param:"+ctx.Param("id"))
}
