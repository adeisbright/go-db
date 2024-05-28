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
	//Initialize Configurations
	config.Init()
	//Use Redis
	redis.Init()
	controller.RunRedis()

	//Use MySQL
	db, conn, err := mysql.MySQLConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	mysql.Migration(db)
	controller.InsertMySQLRecords(db)
}
