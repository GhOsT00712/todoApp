package todo

import (
	"errors"
	"math/rand"
	"time"
)

type TaskManager struct {
	tasks map[int]Task
}

var tm *TaskManager

func GetInstance() *TaskManager {
	if tm == nil {
		tm = &TaskManager{
			tasks: make(map[int]Task),
		}
	}
	return tm
}

// Method to add a Task
// This function adds a new task to the tasks map.
// It validates the input, sets default values, and assigns a random ID.
// Returns an error if the task is invalid (empty header or description).
func (tm *TaskManager) CreateTask(newTask *Task) error {
	if newTask.Description == "" || newTask.Header == "" {
		return errors.New("invalid task")
	}

	//by ref
	// newTask.Id = rand.Intn(1000)
	// newTask.CreatedAt = time.Now()
	// newTask.IsCompleted = false
	// newTask.CompletedAt = time.Time{}
	// newTask.DueDate = time.Time{}

	//create new task state
	task := &Task{
		Id:          rand.Intn(1000),
		Description: newTask.Description,
		Header:      newTask.Header,
		CreatedAt:   time.Now(),
		Priority:    newTask.Priority,
		IsCompleted: false,
		CompletedAt: time.Time{},
		DueDate:     time.Time{},
	}

	// tm.tasks[newTask.Id] = *newTask
	tm.tasks[task.Id] = *task
	return nil
}

// Method to delete a Task
// This function delete a new task to the tasks map.
// Returns an error if the task id is invalid.
func (tm *TaskManager) RemoveTask(id int) error {
	_, exists := tm.tasks[id]

	if !exists {
		return errors.New("task not found")
	}
	delete(tm.tasks, id)
	return nil
}

// MarkTaskCompleted marks a task as completed by its ID
func (tm *TaskManager) MarkTaskCompleted(id int) error {
	task, exists := tm.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	task.IsCompleted = true
	task.CompletedAt = time.Now()
	tm.tasks[id] = task
	return nil
}

// Method to print the Task
func (tm *TaskManager) ToString() string {
	var result string
	for _, task := range tm.tasks {
		result += task.toString() + "\n"
	}
	return result
}
