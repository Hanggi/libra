package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Context struct {
	//	xContext
	w http.ResponseWriter
	r *http.Request
}

type xContext = context.Context

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的

	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

var middlewares []Controller

/// vvvvvvvvvvvvvvvvvvvvvvvvv
type Handler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(rw, r, next)
}

type middleware struct {
	handler http.Handler
	next    *middleware
}

// middleware interface imple
func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ctx Context
	ctx.r = r
	ctx.w = w
	//	for _, value := range middlewares {
	//		//		value(ctx)
	//	}
	m.handler.ServeHTTP(w, r)
}

type Controller func(Context, next http.HandlerFunc)

func Wrap(handler http.Handler) Handler {
	return HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		handler.ServeHTTP(rw, r)
		next(rw, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	//	router := http.NewServeMux()
	//	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//		fmt.Println("vvvv\n")
	//		fmt.Fprintf(w, "hello!!!")
	//	})

	server := http.Server{
		Addr: "127.0.0.1:9090",
		//		Handler: &middleware{Libra.Router},
		//		Handler: Wrap(router),
	}

	http.HandleFunc("/", Wrap(hello))

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error eeee")
	}
}
