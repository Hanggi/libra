package libra

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// "../../libra"

// Log vv
type Log struct {
}

// Logger ...
func Logger() func(*Context) {
	return func(ctx *Context) {
		defer calcTimeEnd(time.Now(), func(d time.Duration) {
			log.WithFields(log.Fields{
				// "Method": ctx.,
				"time": d,
			}).Info("[!!" + ctx.Method + "] " + ctx.URL.Path)
		})
		ctx.Next()
	}

}

func calcTimeEnd(start time.Time, callback func(duration time.Duration)) {
	d := time.Since(start)
	callback(d)
}
