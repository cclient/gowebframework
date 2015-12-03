package tool

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var conn redis.Conn

func init() {

}

//func Test_Redis(t *testing.T) {
//	conn, err := Dial()
//	reply, err := conn.Do("hset", "warnmatch", "new key", 1)
//	fmt.Println(reply)
//	fmt.Println(err)
//	//	if err != nil {
//	//		conn.Close()
//	//		return nil, err
//	//	}
//	//	return conn, err
//}

//func Insert()

//
func Do(commandName string, args ...interface{}) {
	conn, err := Dial()
	if err != nil {
		conn.Do("hset", args)
		conn.Close()
	}
}
func Dial() (redis.Conn, error) {
	c, err := redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		return nil, err
	}
	_, err = c.Do("SELECT", "4")
	if err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}
