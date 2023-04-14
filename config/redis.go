package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnRedis() {
	addr := fmt.Sprintf("%s:%s", GetEnv("REDIS_HOST"), GetEnv("REDIS_PORT"))
	redisDB, _ := strconv.Atoi(GetEnv("REDIS_DB"))
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: GetEnv("REDIS_PASSWORD"),
		DB:       redisDB,
	})
	ctx := context.Background()
	if err := Redis.Ping(ctx).Err(); err != nil {
		fmt.Println("redis: can't ping to redis")
		os.Exit(1)
	}
	fmt.Println("redis: connection opened to redis")
}
