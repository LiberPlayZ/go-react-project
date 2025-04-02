package queries

var GetAllUsersQuery = "SELECT id, username, password, role, created_at, updated_at FROM users"

var CreateUserQuery = `INSERT INTO users (id , username , password , role) VALUES ($1, $2, $3, $4)`

var GetUserByIdQuery = `SELECT id , username, password , role , created_at , updated_at FROM users WHERE id = $1 `
