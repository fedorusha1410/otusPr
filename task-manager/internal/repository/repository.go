package repository

import (
	"database/sql"
	"fmt"
	"task-manager/internal/model/task"
	"time"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) GetTasks(userId int, role string, filter *task.TaskFilter) ([]*task.Task, error) {
	query := `
		SELECT id, title, status, note, created_time, updated_time, priority, author_id
		FROM tasks
		WHERE 1=1
	`
	args := []interface{}{}
	argId := 1

	if role != "1" {
		query += fmt.Sprintf(" AND author_id = $%d", argId)
		args = append(args, userId)
		argId++
	}

	if filter != nil && filter.Status != "" {
		query += fmt.Sprintf(" AND status = $%d", argId)
		args = append(args, filter.Status)
		argId++
	}

	if filter != nil && !filter.CreatedAfter.IsZero() {
		query += fmt.Sprintf(" AND created_time >= $%d", argId)
		args = append(args, filter.CreatedAfter)
		argId++
	}
	if filter != nil && !filter.CreatedBefore.IsZero() {
		query += fmt.Sprintf(" AND created_time <= $%d", argId)
		args = append(args, filter.CreatedBefore)
		argId++
	}

	if filter != nil && !filter.UpdatedAfter.IsZero() {
		query += fmt.Sprintf(" AND updated_time >= $%d", argId)
		args = append(args, filter.UpdatedAfter)
		argId++
	}
	if filter != nil && !filter.UpdatedBefore.IsZero() {
		query += fmt.Sprintf(" AND updated_time <= $%d", argId)
		args = append(args, filter.UpdatedBefore)
		argId++
	}

	rows, err := repository.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task
	for rows.Next() {
		t := &task.Task{}
		var updatedTime sql.NullTime
		var createdTime time.Time

		err := rows.Scan(
			&t.Id, &t.Title, &t.Status, &t.Note,
			&createdTime, &updatedTime, &t.Priority, userId,
		)
		if err != nil {
			return nil, err
		}
		t.CreatedTime = createdTime
		if updatedTime.Valid {
			t.UpdatedTime = updatedTime.Time
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

func (repository *Repository) GetTaskById(id int) (*task.Task, error) {

	t := &task.Task{}
	var updatedTime sql.NullTime
	err := repository.db.QueryRow("SELECT id, title, note, priority, status, updated_time FROM tasks WHERE id = $1", id).
		Scan(&t.Id, &t.Title, &t.Note, &t.Priority, &t.Status, &updatedTime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if updatedTime.Valid {
		t.UpdatedTime = updatedTime.Time
	}
	return t, nil
}

func (repository *Repository) UpdateTask(id int, newData *task.Task) error {

	query := `UPDATE tasks SET title=$1, note=$2, priority=$3, status=$4, updated_time=$5 WHERE id=$6`
	res, err := repository.db.Exec(query, newData.Title, newData.Note, newData.Priority, newData.Status, newData.UpdatedTime, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (repository *Repository) DeleteTask(id int) error {

	res, err := repository.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (repository *Repository) Save(task *task.Task) error {
	query := `INSERT INTO tasks (title, note, priority, status, updated_time) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := repository.db.QueryRow(query, task.Title, task.Note, task.Priority, task.Status, task.UpdatedTime).Scan(&task.Id)
	return err
}
