package libra

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// "../../libra"

// CalcTimeEnd vv
func CalcTimeEnd(start time.Time, callback func(duration time.Duration)) {
	d := time.Since(start)
	callback(d)
}

type FormUtil struct {
}

func (f FormUtil) IsNumber(str string) bool {
	getint, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		fmt.Println(err)
		return false
	}

	if getint > 100 {
		//太大了
	}

	return true
}

func (f FormUtil) IsHan(str string) bool {
	if m, err := regexp.MatchString("^\\p{Han}+$", str); !m {
		if err != nil {
			fmt.Println(err)
			return false
		}
		return false
	}

	return true
}

func (f FormUtil) IsEng(str string) bool {
	if m, err := regexp.MatchString("^[a-zA-Z]+$", str); !m {
		if err != nil {
			fmt.Println(err)
			return false
		}
		return false
	}

	return true
}

func (f FormUtil) IsEmail(str string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, str); !m {
		return false
	}

	return true
}

// Convenience function for printing to stdout
func PP(a ...interface{}) {
	fmt.Println(a)
}

func P(v interface{}) {
	fmt.Println(v)
}

func V(str string, v interface{}) {
	fmt.Printf(str+": %v \n", v)
}

func Version() string {
	return "0.1"
}

func DebugP(str string) {
	if DEBUG {
		fmt.Println(str)
	}
}

//func loadConfig() {
//	file, err := os.Open("config.json")
//	if err != nil {
//		log.Fatalln("Cannot open config file", err)
//	}
//	decoder := json.NewDecoder(file)
//	config = Configuration{}
//	err = decoder.Decode(&config)
//	if err != nil {
//		log.Fatalln("Cannot get configuration from file", err)
//	}
//}
