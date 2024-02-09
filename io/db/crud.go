package db

import (
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

///////////////////
// Create, Read, Update, Delete operations on the tabe task
///////////////////

// CreateTask creates a new task in the database.
func (db *DB) CreateTask(task Task) error {
	_, err := db.conn.Exec("INSERT INTO task (duration, task, category) VALUES (?, ?, ?)", task.Duration, task.Task, task.Category)
	return err
}

// ReadTask reads a task from the database.
func (db *DB) ReadTask(id uuid.UUID) (Task, error) {
	var task Task
	err := db.conn.QueryRow("SELECT * FROM task WHERE id = ?", id).Scan(&task.ID, &task.Duration, &task.Task, &task.Category)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

// UpdateTask updates a task in the database.
func (db *DB) UpdateTask(task Task) error {
	_, err := db.conn.Exec("UPDATE task SET duration = ?, task = ?, category = ? WHERE id = ?", task.Duration, task.Task, task.Category, task.ID)
	return err
}

// DeleteTask deletes a task from the database.
func (db *DB) DeleteTask(id uuid.UUID) error {
	_, err := db.conn.Exec("DELETE FROM task WHERE id = ?", id)
	return err
}

// ListTasks retrieves the last 100 tasks from the database.
func (db *DB) ListTasks() ([]Task, error) {
	rows, err := db.conn.Query("SELECT * FROM task ORDER BY duration DESC LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Duration, &task.Task, &task.Category); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
