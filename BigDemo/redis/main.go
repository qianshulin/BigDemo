package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "43.139.139.207:6379",
		Password: "123456",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("连接失败")
	}
	fmt.Println("连接成功")

	set := rdb.Set("tes3", "damo", 0)
	fmt.Println(set)

}
