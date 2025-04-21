package queries

var GetAllTodosQuery = "SELECT id, title, description, completed, userid, created_at, updated_at FROM todos ORDER BY created_at ASC"

var CreateTodoQuery = `INSERT INTO todos (id , title , description , completed , userid) VALUES ($1, $2, $3, $4, $5)
					   RETURNING id, title, description, completed, userid, created_at, updated_at`

var UpdateTodoToCompletedQuery = `UPDATE todos
		SET completed = true , updated_at = Now()
		WHERE id = $1;
		`

var DeleteTodoQuery = `DELETE FROM todos WHERE id = $1`

var GetUserTodosQuery = `
SELECT todos.id, todos.title, todos.description, todos.completed, todos.created_at, todos.updated_at FROM users 
JOIN todos ON todos.id = ANY(users.todos) WHERE users.id = $1
`

var AddTodoToUserQuery = `UPDATE users SET todos = array_append(todos,$1) WHERE id = $2`
