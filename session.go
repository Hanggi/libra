package libra

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/garyburd/redigo/redis"
)

// "../../libra"

// Session vv
type Session struct {
	sessionID string
	lifeTime  int64
	// lock      sync.Mutex
}

// New vv
func (s Session) New() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Connect to Redis server failed!")
		return
	}
	if _, err := c.Do("AUTH", "qq110119120"); err != nil {
		fmt.Println("Auth failed!")
		c.Close()
	}
	v, err := c.Do("PING")
	fmt.Println(v)

	fmt.Println(s.NewSessionID())

	// v2, err := redis.String(c.Do("GET", "name"))
	// fmt.Println(v2)
}

// NewSessionID vv
func (s Session) NewSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		// return ""
		fmt.Println(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
