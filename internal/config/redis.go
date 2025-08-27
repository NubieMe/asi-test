package config

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	redisstorage "github.com/gofiber/storage/redis/v3"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (*redis.Client, *session.Store) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatalf("Cannot connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")

	store := session.New(session.Config{
		Storage: redisstorage.New(redisstorage.Config{
			Host: "localhost",
			Port: 6379,
		}),
	})

	return redisClient, store
}
