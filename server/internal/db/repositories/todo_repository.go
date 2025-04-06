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

func (r *TodoRepository) CreateTodo(todo models.Todo) (*models.Todo, error) {
	var resultTodo models.Todo
	err := r.DB.QueryRow(queries.CreateTodoQuery, todo.ID, todo.Title, todo.Description, false).Scan(
		&resultTodo.ID,
		&resultTodo.Title,
		&resultTodo.Description,
		&resultTodo.Completed,
		&resultTodo.CreatedAt,
		&resultTodo.UpdatedAt,
	)
	if err != nil {
		log.Println("Error inserting a todo", err)
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
