package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	"crm-backend/internal/rybakcrm/config"
)

func NewRedisDb(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DB.Redis.Host, cfg.DB.Redis.Port),
		Password: cfg.DB.Redis.Password,
		DB:       cfg.DB.Redis.Db,
	})

	return rdb
}
