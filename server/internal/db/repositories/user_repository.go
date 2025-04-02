package repositories

import (
	"database/sql"
	"server/internal/db/queries"
	"server/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {

	rows, err := r.DB.Query(queries.GetAllUsersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) CreateUser(user models.User) error {
	_, err := r.DB.Exec(queries.CreateUserQuery, user.ID, user.Username, user.Password, user.Role)
	return err
}

func (r *UserRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow(queries.GetUserByIdQuery, id).Scan(
		&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no user found
		}
		return nil, err
	}

	return &user, nil
}
