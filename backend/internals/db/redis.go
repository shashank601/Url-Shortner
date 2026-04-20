package db


import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {

    addr := os.Getenv("REDIS_ADDR")
    if addr == "" {
        addr = "localhost:6379"
    }

    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: os.Getenv("REDIS_PASSWORD"),
        DB:       parseIntEnv("REDIS_DB", 0),
    })

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := client.Ping(ctx).Result()
    if err != nil {
        log.Fatal("could not connect to redis:", err)
    }

    log.Println("connected to redis")
    return client
}

func parseIntEnv(key string, fallback int) int {
    v := os.Getenv(key)
    if v == "" {
        return fallback
    }

    var n int
    _, err := fmt.Sscanf(v, "%d", &n)
    if err != nil {
        return fallback
    }
    return n
}