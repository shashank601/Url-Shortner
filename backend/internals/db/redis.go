package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {

	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))

	client := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("could not connect to redis:", err)
	}

	log.Println("connected to redis")
	return client
}
