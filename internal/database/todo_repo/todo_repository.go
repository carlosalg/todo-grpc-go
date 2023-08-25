package todorepo

import (
	"database/sql"
	"fmt"

	"github.com/carlosalg/todo-grpc-go/internal/database/db"
)

type TodoRepository struct {
	db *sql.DB
}

type Todo struct {
	ID        int32
	Title     string
	Completed bool
}

func NewTodoRepository() (*TodoRepository, error) {
	db.InitDB()

	return &TodoRepository{
		db: db.DB,
	}, nil

}

func (r *TodoRepository) CreateTodo(todo Todo) error {
	query := "INSERT INTO todos (title, completed) VALUES (?, ?)"
	_, err := r.db.Exec(query, todo.Title, todo.Completed)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) GetTodoByID(id int32) (Todo, error) {
	query := "SELECT id, title, completed FROM todos WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var todo Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return todo, fmt.Errorf("GetTodoByID %d: no Todo ", id)
		}
		return todo, fmt.Errorf("GetTodoByID %d: %v", id, err)
	}
	return todo, nil
}

func (r *TodoRepository) DeleteTodo(id int32) error {
	query := "DELETE FROM todos WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("DeleteTodo %d: no Todo found", id)
	}

	return nil
}
