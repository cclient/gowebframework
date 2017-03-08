package tool

import (
	"gopkg.in/redis.v4"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     50,
		PoolTimeout:  30 * time.Second,
	})
	//clearh all key
	//client.FlushDb()
}

func getNewClient() *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     50,
		PoolTimeout:  30 * time.Second,
	})
	return client
}

func isConnection() bool {
	reply := client.Ping().String()
	if reply == "ping: PONG" {
		return true
	}
	//"ping: dial tcp :6379: getsockopt: connection refused"
	//dial
	return false
}

func GetClient() *redis.Client {
	if !isConnection() {
		client = getNewClient()
	}
	return client
}

func LRange(client *redis.Client, key string, start int64, end int64) ([]string, error) {
	return client.LRange(key, start, end).Result()
}

func LTrim(client *redis.Client, key string, start int64, end int64) (string, error) {
	return client.LTrim(key, start, end).Result()
}

func LPushArr(client *redis.Client, key string, args []interface{}) (int64, error) {
	return client.LPush(key, args...).Result()
}

func RPushArr(client *redis.Client, key string, args []interface{}) (int64, error) {
	return client.RPush(key, args...).Result()
}

func LPush(client *redis.Client, key string, item interface{}) (int64, error) {
	return client.LPush(key, item).Result()
}

func RPush(client *redis.Client, key string, item interface{}) (int64, error) {
	return client.RPush(key, item).Result()
}
