package todo

import "testing"

func TestNewTask(t *testing.T) {
	// Test case 1: Create a valid task
	t.Run("Create Valid Task", func(t *testing.T) {
		tm := GetInstance()
		newTask := &Task{
			Header:      "Test Task",
			Description: "This is a test task",
			Priority:    MEDIUM,
		}
		err := tm.CreateTask(newTask)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check if the task was added to the map
		if len(tm.tasks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tm.tasks))
		}
	})

	// Test case 2: Create an invalid task (empty header)
	t.Run("Create Invalid Task - Empty Header", func(t *testing.T) {
		tm := GetInstance()
		newTask := &Task{
			Header:      "",
			Description: "This is an invalid task",
			Priority:    LOW,
		}
		err := tm.CreateTask(newTask)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		if err.Error() != "invalid task" {
			t.Errorf("Expected error message 'invalid task', got '%s'", err.Error())
		}
	})

	// Test case 3: Create an invalid task (empty description)
	t.Run("Create Invalid Task - Empty Description", func(t *testing.T) {
		tm := GetInstance()
		newTask := &Task{
			Header:      "Invalid Task",
			Description: "",
			Priority:    HIGH,
		}
		err := tm.CreateTask(newTask)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		if err.Error() != "invalid task" {
			t.Errorf("Expected error message 'invalid task', got '%s'", err.Error())
		}
	})

}

func TestDeleteTask(t *testing.T) {
	// Test case 1: Delete an existing task
	t.Run("Delete Existing Task", func(t *testing.T) {
		tm := GetInstance()
		// Clear tasks before test
		tm.tasks = make(map[int]Task)

		// Create a task
		newTask := &Task{
			Header:      "Test Task",
			Description: "This is a test task",
			Priority:    MEDIUM,
		}
		err := tm.CreateTask(newTask)
		if err != nil {
			t.Fatalf("Failed to create task: %v", err)
		}

		// Get the ID of the created task
		var taskID int
		for id := range tm.tasks {
			taskID = id
			break
		}

		// Delete the task
		err = tm.RemoveTask(taskID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check if the task was removed
		if len(tm.tasks) != 0 {
			t.Errorf("Expected 0 tasks, got %d", len(tm.tasks))
		}
	})

	// Test case 2: Delete a non-existing task
	t.Run("Delete Non-Existing Task", func(t *testing.T) {
		tm := GetInstance()
		// Clear tasks before test
		tm.tasks = make(map[int]Task)

		// Try to delete a non-existing task
		err := tm.RemoveTask(999)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		if err.Error() != "task not found" {
			t.Errorf("Expected error message 'task not found', got '%s'", err.Error())
		}
	})
}

func TestMarkTaskCompleted(t *testing.T) {
	// Test case 1: Mark an existing task as completed
	t.Run("Mark Existing Task as Completed", func(t *testing.T) {
		tm := GetInstance()
		// Clear tasks before test
		tm.tasks = make(map[int]Task)

		// Create a task
		newTask := &Task{
			Header:      "Test Task",
			Description: "This is a test task",
			Priority:    MEDIUM,
		}
		err := tm.CreateTask(newTask)
		if err != nil {
			t.Fatalf("Failed to create task: %v", err)
		}

		// Get the ID of the created task
		var taskID int
		for id := range tm.tasks {
			taskID = id
			break
		}

		// Mark the task as completed
		err = tm.MarkTaskCompleted(taskID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check if the task was marked as completed
		task, exists := tm.tasks[taskID]
		if !exists {
			t.Errorf("Task with ID %d not found", taskID)
		}
		if !task.IsCompleted {
			t.Errorf("Expected task to be completed, but it wasn't")
		}
		if task.CompletedAt.IsZero() {
			t.Errorf("Expected CompletedAt to be set, but it wasn't")
		}
	})

	// Test case 2: Try to mark a non-existing task as completed
	t.Run("Mark Non-Existing Task as Completed", func(t *testing.T) {
		tm := GetInstance()
		// Clear tasks before test
		tm.tasks = make(map[int]Task)

		// Try to mark a non-existing task as completed
		err := tm.MarkTaskCompleted(999)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		if err.Error() != "task not found" {
			t.Errorf("Expected error message 'task not found', got '%s'", err.Error())
		}
	})
}
