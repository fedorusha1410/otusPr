package repository

import (
	"database/sql"
	"task-manager/internal/model/task"
	"time"
)

const taskFile = "tasks.json"

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) GetTasks() ([]*task.Task, error) {
	rows, err := repository.db.Query("SELECT id, title, note, priority, status, updated_time FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task
	for rows.Next() {
		t := &task.Task{}
		var updatedTime sql.NullTime
		err := rows.Scan(&t.Id, &t.Title, &t.Note, &t.Priority, &t.Status, &updatedTime)
		if err != nil {
			return nil, err
		}
		if updatedTime.Valid {
			t.UpdatedTime = updatedTime.Time
		} else {
			t.UpdatedTime = time.Time{}
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

func (repository *Repository) Save(task task.Task) error {
	query := `INSERT INTO tasks (title, note, priority, status, updated_time) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := repository.db.QueryRow(query, task.Title, task.Note, task.Priority, task.Status, task.UpdatedTime).Scan(&task.Id)
	return err
}
