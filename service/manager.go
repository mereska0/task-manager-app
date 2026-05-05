package service

import (
	"context"
	"database/sql"
	"errors"
	"tmanager-app/model"
)

var ErrNotFound = errors.New("not found")

type TaskService struct {
	db *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) AddTask(ctx context.Context, text string) (model.Task, error) {
	var task model.Task

	err := s.db.QueryRowContext(ctx,
		"INSERT INTO tasks (text, isdone) VALUES ($1, $2) RETURNING id, text, isdone",
		text, false,
	).Scan(&task.ID, &task.Text, &task.IsDone)
	return task, err
}

func (s *TaskService) GetTasks(ctx context.Context) ([]model.Task, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		rows.Scan(&task.ID, &task.Text, &task.IsDone)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *TaskService) UpdateTaskText(ctx context.Context, text string, id int) error {
	res, err := s.db.ExecContext(ctx, "UPDATE tasks SET text=$1 WHERE id=$2", text, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}
func (s *TaskService) UpdateTaskStatus(ctx context.Context, isDone bool, id int) error {
	res, err := s.db.ExecContext(ctx, "UPDATE tasks SET isdone=$1 WHERE id=$2", isDone, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	res, err := s.db.ExecContext(ctx, "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *TaskService) ClearTasks(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM TASKS")
	if err != nil {
		return err
	}
	return nil
}
