package todo

import "testing"

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

// func TestDeleteTask(t *testing.T) {
// }

// func TestMarkTaskCompleted(t *testing.T) {
// }
