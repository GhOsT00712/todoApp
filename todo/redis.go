package todo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

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
	ctx := context.Background()
	_, err := rc.redisClient.HSet(ctx, key, map[string]interface{}{
		"id":          task.Id,
		"header":      task.Header,
		"description": task.Description,
		"isCompleted": task.IsCompleted,
		"createdAt":   task.CreatedAt,
		"CompletedAt": task.CompletedAt,
		"dueDate":     task.DueDate,
		"priority":    task.Priority.String(),
	}).Result()
	if err != nil {
		return "", err
	}

	return key, nil
}

func (rc *Redis) ScanTask() ([]string, error) {
	res, _, err := rc.redisClient.Scan(context.Background(), 0, "Task:*", 1000).Result()
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (rc *Redis) RemoveTask(id string) error {
	ctx := context.Background()
	_, err := rc.redisClient.Del(ctx, id).Result()
	if err != nil {
		return err
	}
	return nil
}

func (rc *Redis) GetTask(id string) (task Task, err error) {
	ctx := context.Background()
	res, err := rc.redisClient.HGetAll(ctx, id).Result()
	task.Id, _ = strconv.Atoi(res["id"])
	task.Header = res["header"]
	task.Description = res["description"]
	task.CompletedAt, _ = time.Parse(time.RFC3339, res["completedAt"])
	task.CreatedAt, _ = time.Parse(time.RFC3339, res["createdAt"])
	task.DueDate, _ = time.Parse(time.RFC3339, res["dueDate"])
	priorityValue, _ := strconv.Atoi(res["priority"])
	task.Priority = Priority(priorityValue)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
