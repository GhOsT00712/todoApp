package todo

import (
	"fmt"
	"time"
)

// Data structure to hold Task data
type Task struct {
	Id          int       `redis:"id"`
	Header      string    `redis:"header"`
	Description string    `redis:"description"`
	IsCompleted bool      `redis:"isCompleted"`
	CreatedAt   time.Time `redis:"createdAt"`
	CompletedAt time.Time `redis:"completedAt"`
	DueDate     time.Time `redis:"dueDate"`
	Priority    Priority  `redis:"priority"`
}

// Method to print the Task
func (task *Task) ToString() string {
	return fmt.Sprintf("Task ID: %d\nHeader: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\nCompleted At: %s\nDue: %s\nPriority: %d",
		task.Id,
		task.Header,
		task.Description,
		task.IsCompleted,
		task.CreatedAt.Format(time.RFC3339),
		task.CompletedAt.Format(time.RFC3339),
		task.DueDate.Format(time.RFC3339),
		task.Priority)
}
