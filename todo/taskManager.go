package todo

import (
	"errors"
	"math/rand"
	"time"
)

type TaskManager struct {
	dbClient *Redis
}

var tm *TaskManager

func GetInstance() *TaskManager {
	if tm == nil {
		tm = &TaskManager{
			dbClient: GetClient(),
		}
	}
	return tm
}

func (tm *TaskManager) CreateTask(newTask *Task) (string, error) {
	if newTask.Description == "" || newTask.Header == "" {
		return "", errors.New("invalid task")
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
	// tm.tasks[task.Id] = *task
	id, err := tm.dbClient.AddTask(task)
	if err == nil {
		return id, nil
	}

	return "", err
}

func (tm *TaskManager) GetAllTasks() ([]string, error) {
	res, err := tm.dbClient.ScanTask()
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		println(v)
	}

	return res, nil
}

// Method to delete a Task
func (tm *TaskManager) RemoveTask(id string) error {
	err := tm.dbClient.RemoveTask(id)
	if err != nil {
		return err
	}
	return nil
}

// // MarkTaskCompleted marks a task as completed by its ID
func (tm *TaskManager) MarkTaskCompleted(id string) error {
	task, err := tm.dbClient.GetTask(id)
	if err != nil {
		return err
	}
	task.IsCompleted = true
	task.CompletedAt = time.Now()
	_, err = tm.dbClient.AddTask(&task)
	if err != nil {
		return err
	}
	return nil
}

// Method to print the Task
// func (tm *TaskManager) ToString() string {
// 	var result string
// 	for _, task := range tm.tasks {
// 		result += task.toString() + "\n"
// 	}
// 	return result
// }
