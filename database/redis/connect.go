package redis

import (
	"github.com/go-redis/redis/v8"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

var DB *redis.Client

func Connect() {
	DB = redis.NewClient(&redis.Options{
		Addr:     config.Redis_Addr,
		Password: config.Redis_Pass,
		DB:       0,
	})
}
