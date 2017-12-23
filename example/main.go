package main

import (
	"fmt"
	"net/http"

	"github.com/Hanggi/libra"
	"github.com/Hanggi/libra/example/route"

	_ "github.com/go-sql-driver/mysql"
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

func middleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("This is a 111 middleware!")

	//	defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
	//		fmt.Println(d, "\n")
	//	})
	next(rw, r)
	fmt.Println("middleware 111 end!!!")
}
func middleware2(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("This is a middleware22222!")

	//	defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
	//		fmt.Println(d, "\n")
	//	})
	next(rw, r)
	fmt.Println("This is a middleware22222  end!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := libra.New()
	port := app.Config.Port
	port = 5555

	router := app.LRouter

	app.Static("/public", "public")

	// db, err := sql.Open("mysql", "root:qq110119120@/fortest")
	// checkErr(err)

	// stmt, err := db.Prepare("INSERT test1 SET username=?,age=?,salary=?")
	// checkErr(err)

	// res, err := stmt.Exec("ff", "235", "1129")1
	// checkErr(err)

	// id, err := res.LastInsertId()

	// fmt.Println(id)

	router.Use(middleware)
	router.Use(middleware2)

	router.GET("/form/input", route.FormInput)
	router.POST("/form/input", route.FormInputPost)

	router.GET("/", route.Index)
	router.GET("/test", route.Test)
	router.GET("/test/:id", route.Test)
	router.GET("/test/:id/*action", route.Test)
	router.GET("/post", func(ctx libra.Context) {
		fmt.Println("in get post")

		ctx.Render("post", data{"data's name", 12, "mamada"})
	})

	app.ListenAnd(port, func() {
		fmt.Printf("listening at port: %d\n", port)
	})
}
