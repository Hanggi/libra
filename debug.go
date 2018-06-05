package libra

import "log"

// "../../libra"

// Log vv
type LibraDebug struct {
	DebugPrint string
	IsShow     bool
}

func init() {
	log.SetPrefix("[Libra]")
	log.SetFlags(log.Lshortfile | log.Ldate)
	//LibraDebug.IsShow = false
}

func (debug *LibraDebug) Print(v ...interface{}) {
	log.Println(v)
}

func (debug *LibraDebug) Hide() {
	debug.IsShow = false
}
