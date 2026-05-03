package service

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreate(t *testing.T) {
	db := setupTestDB(t)
	service := NewTaskService(db)
	task, err := service.AddTask(context.Background(), "milk")
	if err != nil {
		t.Fatal(err)
	}

	if task.Text != "milk" {
		t.Errorf("expected milk, got %s", task.Text)
	}

	if task.ID == 0 {
		t.Errorf("expected ID to be set")
	}
}

func TestRead(t *testing.T) {
	db := setupTestDB(t)
	service := NewTaskService(db)
	_, err := service.AddTask(context.Background(), "banana")
	if err != nil {
		t.Fatal(err)
	}
	tasks, err := service.GetTasks(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)
	service := NewTaskService(db)
	task, err := service.AddTask(context.Background(), "cheese")
	if err != nil {
		t.Fatal(err)
	}
	err = service.UpdateTaskText(context.Background(), "banana", task.ID)
	if err != nil {
		t.Fatal(err)
	}
	tasks, err := service.GetTasks(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if tasks[0].Text != "banana" {
		t.Errorf("expected text \"banana\", got \"%s\"", task.Text)
	}
	err = service.UpdateTaskStatus(context.Background(), true, task.ID)

	tasks, err = service.GetTasks(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !tasks[0].IsDone {
		t.Error("expected status \"true\", got \"false\"")
	}
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)
	service := NewTaskService(db)
	task, err := service.AddTask(context.Background(), "chocolate")
	if err != nil {
		t.Fatal(err)
	}
	err = service.DeleteTask(context.Background(), task.ID)
	if err != nil {
		t.Fatal(err)
	}
	tasks, err := service.GetTasks(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 0 {
		t.Errorf("no tasks expected, got %s", tasks[0].Text)
	}
}
