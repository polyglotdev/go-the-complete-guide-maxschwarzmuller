package main

import (
	"fmt"
	"log"
	"time"
)

type Task struct {
	ID        int
	Name      string
	Completed bool
	CreatedAt time.Time
}

func (t *Task) MarkComplete() {
	t.Completed = true
}

func (t *Task) Details() string {
	return fmt.Sprintf("ID: %d, Name: %s, Completed: %t, CreatedAt: %v", t.ID, t.Name, t.Completed, t.CreatedAt)
}

type TaskManager struct {
	tasks []*Task
}

func (tm *TaskManager) Add(task *Task) {
	tm.tasks = append(tm.tasks, task)
}

func (tm *TaskManager) Complete(taskID int) error {
	for _, task := range tm.tasks {
		if task.ID == taskID {
			task.MarkComplete()
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", taskID)
}

func (tm *TaskManager) List() []*Task {
	return tm.tasks
}

func main() {
	task1 := &Task{ID: 1, Name: "Buy milk", Completed: false, CreatedAt: time.Now()}
	task2 := &Task{ID: 2, Name: "Read book", Completed: false, CreatedAt: time.Now()}

	manager := TaskManager{}
	manager.Add(task1)
	manager.Add(task2)
	if err := manager.Complete(2); err != nil {
		log.Fatalf("Error completing task: %v", err)
	}

	if err := manager.Complete(1); err != nil {
		log.Fatalf("Error completing task: %v", err)
	}

	for _, task := range manager.List() {
		fmt.Println(task.Details())
	}
}
