// To test and use the dbs
package controller

import (
	"fmt"

	"github.com/adeisbright/go-db/redis"
)

func Run() {
	redisCacheKey := "test:user"
	nameValue := "Adeleke Bright"
	cmd, _ := redis.NewRedisFactory().GetRedisValue(redisCacheKey)
	name, err := cmd.Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	if name == "" {
		name = nameValue
		//If I want to set more than one values to more than one keys, I can leverage
		// concurrency here
		if err := redis.NewRedisFactory().SetRedisValue(redisCacheKey, nameValue); err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(name)
}
