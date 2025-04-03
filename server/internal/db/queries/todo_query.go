package queries

var GetAllTodosQuery = "SELECT id, title, description, completed, created_at, updated_at FROM todos ORDER BY created_at ASC"

var CreateTodoQuery = `INSERT INTO todos (id , title , description , completed) VALUES ($1, $2, $3, $4)
					   RETURNING id, title, description, completed, created_at, updated_at`

var UpdateTodoToCompletedQuery = `UPDATE todos
		SET completed = true , updated_at = Now()
		WHERE id = $1;
		`
