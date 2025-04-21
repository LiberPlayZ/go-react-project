package migrations

var Tables = []string{

	// users
	`CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		role VARCHAR(50) NOT NULL,
		todos UUID[],
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`,

	`CREATE TABLE IF NOT EXISTS todos (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(50) UNIQUE NOT NULL,
		description VARCHAR(150)  NOT NULL,
		completed BOOLEAN NOT NULL,
		userid UUID NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`,
}
