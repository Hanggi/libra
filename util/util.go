package util

import "time"

// "../../libra"

func CalcTimeEnd(start time.Time, callback func(duration time.Duration)) {
	d := time.Since(start)
	callback(d)
}
