package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"github.com/shashank601/url-shortner/backend/internals/db"
)

func InitRedis() *redis.Client {
	return db.NewRedisClient()
}
