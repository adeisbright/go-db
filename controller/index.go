package controller

import (
	"fmt"
	"log"

	"github.com/adeisbright/go-db/mysql/entities"
	"github.com/adeisbright/go-db/redis"
	"gorm.io/gorm"
)

func RunRedis() {
	redisCacheKey := "test:user"
	nameValue := "Adeleke Bright"
	rf, err := redis.NewRedisFactory()
	if err != nil {
		panic(err)
	}
	cmd, _ := rf.GetRedisValue(redisCacheKey)
	name, err := cmd.Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	if name == "" {
		name = nameValue
		//If I want to set more than one values to more than one keys, I can leverage
		// concurrency here
		if err := rf.SetRedisValue(redisCacheKey, nameValue); err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(name)
}

type UserSchema struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func InsertMySQLRecords(db *gorm.DB) {
	var user entities.User

	user.Username = "John Dumelo"
	user.Email = "dumelo@gmail.com"
	if result := db.Create(&user); result.Error != nil {
		log.Fatal("Could not create user")
	}
	fmt.Println("User created successfully")
}
