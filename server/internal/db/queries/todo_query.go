package queries

var GetAllTodosQuery = "SELECT id, title, description, completed, created_at, updated_at FROM todos"

var CreateTodoQuery = `INSERT INTO todos (id , title , description , completed) VALUES ($1, $2, $3, $4)`
