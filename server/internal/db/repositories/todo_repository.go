package repositories

import (
	"database/sql"
	"server/internal/db/queries"
	"server/internal/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetAllTodos() ([]models.Todo, error) {

	rows, err := r.DB.Query(queries.GetAllTodosQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) CreateTodo(todo models.Todo) error {
	_, err := r.DB.Exec(queries.CreateTodoQuery, todo.ID, todo.Title, "", false)
	return err
}
