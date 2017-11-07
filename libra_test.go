package libra

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// "./context"
// "github.com/Hanggi/libra/controller"

func TestHTTPListen(t *testing.T) {
	go func() {
		var app = new(App)
		var router = app.LRouter

		router.GET("/test", func(ctx Context) {
			fmt.Fprintf(ctx.w, "hello test!")
		})
		app.Listen(9090)
	}()
	t.Log("waiting 1 second for server startup")
	time.Sleep(1 * time.Second)

	res, err := http.Get("http://127.0.0.1:9090/test")
	if err != nil {
		fmt.Println(err)
	}
	resp, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "hello test!", string(resp[:]))
}
