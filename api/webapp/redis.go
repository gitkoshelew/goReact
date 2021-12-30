package webapp

import (
	"log"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

// ConnectRedis ...
func ConnectRedis(config *Config) *redis.Client {
	addr := config.RedisInfo()
	client = redis.NewClient(&redis.Options{Addr: addr})
	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("Redis address: %s", addr)
		log.Printf("Error trying to ping redis: %v", err)
	}
	return client
}
