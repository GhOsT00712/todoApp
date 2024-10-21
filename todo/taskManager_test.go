package todo

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	t.Run("Create Valid Task", func(t *testing.T) {
		tm := GetInstance()
		newTask := &Task{
			Header:      "Valid Task",
			Description: "This is a valid task",
			Priority:    MEDIUM,
		}
		_, err := tm.CreateTask(newTask)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestRemoveTask(t *testing.T) {
	t.Run("Remove Task", func(t *testing.T) {

		tm := GetInstance()
		newTask := &Task{
			Header:      "Valid Task",
			Description: "This is a valid task",
			Priority:    MEDIUM,
		}
		id, err := tm.CreateTask(newTask)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		err = tm.RemoveTask(id)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestGetAllTask(t *testing.T) {
	t.Run("Get All Tasks", func(t *testing.T) {
		tm := GetInstance()
		_, err := tm.GetAllTasks()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestMarkTaskCompleted(t *testing.T) {
	t.Run("Mark Task Completed", func(t *testing.T) {
		tm := GetInstance()
		newTask := &Task{
			Header:      "Valid Task",
			Description: "This is a valid task",
			Priority:    MEDIUM,
		}
		id, err := tm.CreateTask(newTask)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		err = tm.MarkTaskCompleted(id)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

}
