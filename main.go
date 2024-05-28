package main

import (
	"fmt"

	"github.com/adeisbright/go-db/config"
	"github.com/adeisbright/go-db/controller"
	"github.com/adeisbright/go-db/mysql"
	"github.com/adeisbright/go-db/redis"
)

func main() {
	fmt.Println("Working with DBs in Golang")
	config.Init()
	redis.Init()
	controller.Run()
	mysql.Init()
}
