package main

import (
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Do("PUBLISH", "example", "hello")
	c.Do("PUBLISH", "example", "world")
	c.Do("PUBLISH", "pexample", "foo")
	c.Do("PUBLISH", "pexample", "bar")
}
