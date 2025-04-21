package repositories

import (
	"database/sql"
	"log"
	"server/internal/db/queries"
	"server/internal/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetAllTodos(userId string) ([]models.Todo, error) {
	var Rows *sql.Rows
	if userId == "" {
		rows, err := r.DB.Query(queries.GetAllTodosQuery)
		if err != nil {
			return nil, err
		}
		Rows = rows
	} else {
		rows, err := r.DB.Query(queries.GetUserTodosQuery, userId)
		if err != nil {
			return nil, err
		}
		Rows = rows

	}

	defer Rows.Close()

	var todos []models.Todo
	for Rows.Next() {
		var todo models.Todo
		err := Rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.UserId,
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

func (r *TodoRepository) CreateTodo(todo models.Todo) (*models.Todo, error) {
	var resultTodo models.Todo
	err := r.DB.QueryRow(queries.CreateTodoQuery, todo.ID, todo.Title, todo.Description, false, todo.UserId).Scan(
		&resultTodo.ID,
		&resultTodo.Title,
		&resultTodo.Description,
		&resultTodo.Completed,
		&resultTodo.UserId,
		&resultTodo.CreatedAt,
		&resultTodo.UpdatedAt,
	)
	if err != nil {
		log.Println("Error inserting a todo", err)
		return nil, err
	}

	_, err = r.DB.Exec(queries.AddTodoToUserQuery, todo.ID, todo.UserId)
	if err != nil {
		return nil, err
	}

	return &resultTodo, nil
}

func (r *TodoRepository) UpdateTodoToCompleted(todoId string) error {

	res, err := r.DB.Exec(queries.UpdateTodoToCompletedQuery, todoId)
	if err != nil {
		log.Println("Error updating todo:", err)
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows:", err)
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *TodoRepository) DeleteTodo(todoId string) error {

	res, err := r.DB.Exec(queries.DeleteTodoQuery, todoId)
	if err != nil {
		log.Println("Error deleting todo:", err)
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows:", err)
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
