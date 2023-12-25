package main

import (
	"fmt"
	"init-case/redis"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

/*
* main
activation result :
redis
init 1
init 2

-> This is because main is dependent on redis.
-> init() on redis activates first.
*/
func main() {
	err := redis.Store("foo", "bar")
	_ = err
}
