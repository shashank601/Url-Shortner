package bootstrap

import (
    "context"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {

    opts := &redis.Options{
        Addr:     "localhost:6379", // we will use upstash not local server
        Password: "",               
        DB:       0,                
    }

    client := redis.NewClient(opts)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := client.Ping(ctx).Result()
    if err != nil {
        panic(fmt.Sprintf("Could not connect to Redis: %v", err))
    }

    fmt.Println("Successfully connected to Redis")
    return client
}