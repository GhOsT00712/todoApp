package main

import (
	"fmt"
	"todo"
)

func main() {
	//get TaskManager
	taskManager := todo.GetInstance()

	//Add a task
	newTask := &todo.Task{
		Header:      "Hello Todo",
		Description: "Hello Todo Description",
		Priority:    todo.Priority(todo.HIGH),
	}

	id, err := taskManager.CreateTask(newTask)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}

}
