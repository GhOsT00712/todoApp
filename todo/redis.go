package todo

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var addr = os.Getenv("REDIS_ADDR")
var password = os.Getenv("REDIS_PASSWORD")

type Redis struct {
	redisClient *redis.Client
}

var rc *Redis

func GetClient() *Redis {
	if rc == nil {
		rc = &Redis{
			redisClient: client(),
		}
	}
	return rc
}

func client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // No password set
		DB:       0,        // Use default DB
		Protocol: 2,        // Connection protocol
	})
	return client
}

func (rc *Redis) AddTask(task *Task) (string, error) {
	key := fmt.Sprintf("Task:%d", task.Id)
	println(key)
	_, err := rc.redisClient.HSet(context.Background(), key, map[string]interface{}{
		"id":          task.Id,
		"header":      task.Header,
		"description": task.Description,
		"isCompleted": task.IsCompleted,
		"createdAt":   task.CreatedAt,
		"CompletedAt": task.CompletedAt,
		"dueDate":     task.DueDate,
		"priority":    task.Priority.String(),
	}).Result()
	if err == nil {
		return key, nil
	}

	return "", err
}
