package todo

import (
	"fmt"
	"time"
)

// Data structure to hold Task data
type Task struct {
	Id          int
	Header      string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt time.Time
	DueDate     time.Time
	Priority    Priority
}

// Method to print the Task
func (task *Task) toString() string {
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
