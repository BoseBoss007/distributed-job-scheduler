package shared

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitRedis() {

	redisHost := os.Getenv("REDIS_HOST")

	if redisHost == "" {
		redisHost = "localhost:6379" // default for scaler
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr: redisHost,
	})
}