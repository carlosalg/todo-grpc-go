package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/carlosalg/todo-grpc-go/internal/database/db"
	"github.com/carlosalg/todo-grpc-go/internal/database/todo_repo"
	"github.com/stretchr/testify/assert"
)

// func TestInitDB(t *testing.T) {
// 	db.InitDB()
// 	if db.DB == nil {
// 		t.Errorf("DB not initialized, got nil")
// 	}
// 	db.CloseDB()
// }

func TestCreateTodo(t *testing.T) {
	t.Run("TestCreate", func(t *testing.T) {
		repo, err := todorepo.NewTodoRepository()
		assert.NoError(t, err)

		todo := todorepo.Todo{
			ID:        1,
			Title:     "Test CreateTodo",
			Completed: false,
		}
		err = repo.CreateTodo(todo)
		assert.NoError(t, err)

		createdTodo, err := repo.GetTodoByID(1)
		assert.NoError(t, err)
		assert.Equal(t, todo, createdTodo)
	})
}

func TestGetTodoByID(t *testing.T) {
	t.Run("TestGet", func(t *testing.T) {
		repo, err := todorepo.NewTodoRepository()
		assert.NoError(t, err)

		err = repo.CreateTodo(todorepo.Todo{ID: 2, Title: "Test GetTodo", Completed: false})
		assert.NoError(t, err)

		todo, err := repo.GetTodoByID(2)
		assert.NoError(t, err)

		expectedTodo := todorepo.Todo{ID: 2, Title: "Test GetTodo", Completed: false}
		assert.Equal(t, expectedTodo, todo)
	})
}

func TestDeleteTodo(t *testing.T) {
	t.Run("TestDelete", func(t *testing.T) {
		repo, err := todorepo.NewTodoRepository()
		assert.NoError(t, err)

		todoToDelete := todorepo.Todo{ID: 3, Title: "Test DeleteTodo", Completed: false}
		err = repo.CreateTodo(todoToDelete)
		assert.NoError(t, err)

		err = repo.DeleteTodo(3)
		assert.NoError(t, err)
	})
}

func cleanupTestData() error {
	_, err := db.DB.Exec("DELETE FROM todos WHERE id >= 1")
	if err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	db.InitDB()
	_, err := db.DB.Exec("ALTER TABLE todos AUTO_INCREMENT = 1")
	if err != nil {
		fmt.Printf("Error resetting auto_increment: %v\n", err)
	}

	result := m.Run()

	cleanupErr := cleanupTestData()
	if cleanupErr != nil {
		fmt.Printf("Error cleaning up test data: %v\n", cleanupErr)
	}

	db.CloseDB()
	os.Exit(result)
}
