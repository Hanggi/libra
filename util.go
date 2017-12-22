package libra

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// "../../libra"

// CalcTimeEnd vv
func CalcTimeEnd(start time.Time, callback func(duration time.Duration)) {
	d := time.Since(start)
	callback(d)
}

type formUtil struct {
}

func (f formUtil) IsNumber(str string) bool {
	getint, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if getint > 100 {
		//太大了
	}

	return true
}

func (f formUtil) IsHan(str string) bool {
	if m, err := regexp.MatchString("^\\p{Han}+$", str); !m {
		if err != nil {
			fmt.Println(err)
			return false
		}
		return false
	}

	return true
}

func (f formUtil) IsEng(str string) bool {
	if m, err := regexp.MatchString("^[a-zA-Z]+$", str); !m {
		if err != nil {
			fmt.Println(err)
			return false
		}
		return false
	}

	return true
}

func (f formUtil) IsEmail(str string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, str); !m {
		return false
	}

	return true
}

// Convenience function for printing to stdout
func P(a ...interface{}) {
	fmt.Println(a)
}

func Version() string {
	return "0.1"
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
