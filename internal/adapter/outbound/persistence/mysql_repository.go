package persistence

import (
	"database/sql"

	_"github.com/go-sql-driver/mysql"
	"task-api/internal/domain"
	"task-api/internal/port/outbound"
)

type mysqlRepo struct {
	db *sql.DB
}

func NewMySQLRepo(db *sql.DB) outbound.TaskRepository {
	return &mysqlRepo{db: db}
}

func (r *mysqlRepo) Create(task *domain.Task) error {
	query := "INSERT INTO tasks (title, description) VALUES (?, ?)"
	_, err := r.db.Exec(query, task.Title, task.Description)
	return err
}

func (r *mysqlRepo) GetByID(id int64) (*domain.Task, error) {
	query := "SELECT id, title, description FROM tasks WHERE id = ?"
	row := r.db.QueryRow(query, id)

	task := &domain.Task{}
	err := row.Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *mysqlRepo) GetAll() ([]*domain.Task, error) {
	query := "SELECT id, title, description FROM tasks"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *mysqlRepo) Update(task *domain.Task) error {
	query := "UPDATE tasks SET title = ?, description = ? WHERE id = ?"
	_, err := r.db.Exec(query, task.Title, task.Description, task.ID)
	return err
}

func (r *mysqlRepo) Delete(id int64) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

