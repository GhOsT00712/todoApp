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

	err := taskManager.CreateTask(newTask)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(taskManager.ToString())
	}

	err = taskManager.RemoveTask(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Task removed")
	}

	err = taskManager.RemoveTask(newTask.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Task removed")
	}

}
