package queries

var GetAllUsersQuery = "SELECT id, username, email, password, role, created_at, updated_at FROM users"

var CreateUserQuery = `INSERT INTO users (id , username, email , password , role) VALUES ($1, $2, $3, $4, $5)`

var GetUserByIdQuery = `SELECT id , username, email, password , role , created_at , updated_at FROM users WHERE id = $1 `

var GetUserByEmailQuery = `SELECT id , username, email, password , role , created_at , updated_at FROM users WHERE email = $1 `
