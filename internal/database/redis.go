package database

import (
	"context"
	"fmt"

	"autocomplete/utils/env"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type RedisClient struct {
	Client *redis.Client
	Logger *zerolog.Logger
}

func NewRedisClient(conf env.ConfRedis, logger *zerolog.Logger) (*RedisClient, error) {
	redisHost := conf.Host
	redisPort := conf.Port
	redisPassword := conf.Password
	redisUser := conf.User

	logger.Info().Str("host", redisHost).Int("port", redisPort).Str("user", conf.User).Msg("Creating RedisClient client")

	redisUri := fmt.Sprintf("%s:%d", redisHost, redisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     redisUri,
		Password: redisPassword,
		DB:       0, // use default DB
		Username: redisUser,
	})

	// Test connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Error().Err(err).Str("uri", redisUri).Msg("Failed to connect to RedisClient")
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	logger.Info().Str("uri", redisUri).Msg("RedisClient connection established successfully")

	return &RedisClient{Client: client, Logger: logger}, nil
}
