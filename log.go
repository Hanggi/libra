package libra

import (
	"time"

	"github.com/Hanggi/libra/util"
	log "github.com/sirupsen/logrus"
)

// "../../libra"

// Log vv
type Log struct {
}

// HTTPLogger vv
func (l *Log) HTTPLogger(str string) {
	defer util.CalcTimeEnd(time.Now(), func(d time.Duration) {
		log.WithFields(log.Fields{
			"Method": str,
			"time":   d,
		}).Info("log test")
	})
}

func calcTimeEnd(start time.Time, callback func(duration time.Duration)) {
	d := time.Since(start)
	callback(d)
}
