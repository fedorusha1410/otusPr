package redislog

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type RedisLogger struct {
	client *redis.Client
}

const logTTL = 1 * time.Hour

func NewRedisLogger(addr string) *RedisLogger {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisLogger{
		client: rdb,
	}
}

func (r *RedisLogger) LogAction(ctx context.Context, action string, entity string, details interface{}) error {
	timestamp := time.Now().Format(time.RFC3339Nano)
	key := "log:" + timestamp
	value := map[string]interface{}{
		"action":  action,
		"entity":  entity,
		"details": details,
		"time":    timestamp,
	}
	jsonValue, err := json.Marshal(value)
	if err != nil {
		log.Printf("Logger: Failed to marshal log: %v\n", err)
		return err
	}

	return r.client.Set(key, jsonValue, logTTL).Err()
}

func (r *RedisLogger) LogPrinter(ctx context.Context, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		seen := make(map[string]bool)

		for {
			select {
			case <-ctx.Done():
				log.Println("LogPrinter stopped")
				return
			case <-ticker.C:
				iter := r.client.Scan(0, "log:*", 0).Iterator()

				for iter.Next() {
					key := iter.Val()
					if seen[key] {
						continue
					}

					val, err := r.client.Get(key).Result()
					if err != nil {
						log.Printf("Logger: Error reading key %s: %v\n", key, err)
						continue
					}

					log.Printf("Logger: %s: %s\n", key, val)

					seen[key] = true
				}
				if err := iter.Err(); err != nil {
					log.Printf("Logger: scan error: %v\n", err)
				}
			}
		}
	}()
}
