package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)
/*
#启动redis
redis-server /usr/local/redis-4.0.10/redis.conf
#连接redis
#redis-cli查看帮助
redis-cli -h 127.0.0.1 -p 6379
*/

func main() {
	address := "127.0.0.1:6379"
	c,err := redis.Dial("tcp",address)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	//设置值
	reply,err := c.Do("set","username","何会奇")
	if err != nil {
		panic(err)
	}
	fmt.Println("set-value=",reply)

	//获取值
	reply,err = c.Do("get","username")
	if err != nil {
		panic(err)
	}
	username,err := redis.String(reply,err)
	if err != nil {
		panic(err)
	}

	fmt.Printf("get-value=%s",username)


}
