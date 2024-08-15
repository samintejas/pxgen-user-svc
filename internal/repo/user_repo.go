package repo

import (
	"database/sql"
	"fmt"
	"strings"

	"pxgen.io/user/internal/models"
)

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, error)
	DeleteUser(username string) error
	ExcistsByUsernameAndEmail(username string, email string) (bool, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (uint, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	return []models.User{}, nil
}

func (r *UserRepository) DeleteUser(username string) error {
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
        SELECT id, username, email, password, first_name, last_name, status, created_at, updated_at
        FROM users
        WHERE email = ?
    `
	return r.getUserByQuery(query, email)
}

func (r *UserRepository) GetUserByEmailAndStatus(email string, status string) (*models.User, error) {
	query := `
        SELECT id, username, email, password, first_name, last_name, status, created_at, updated_at
        FROM users
        WHERE email = ?
    `
	return r.getUserByQuery(query, email)
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := `
        SELECT id, username, email, first_name, last_name, status, created_at, updated_at
        FROM users
        WHERE username = ?
    `
	return r.getUserByQuery(query, username)
}

func (r *UserRepository) GetUserByIdAndStatus(email string, status string) (*models.User, error) {
	query := `
        SELECT id, username, email, password, first_name, last_name, status, created_at, updated_at
        FROM users
        WHERE id = ? AND status = ?
    `
	return r.getUserByQuery(query, email, status)
}

func (r *UserRepository) getUserByQuery(query string, args ...any) (*models.User, error) {

	var user models.User
	err := r.db.QueryRow(query, args...).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) (uint, error) {

	query := "INSERT INTO users (username,first_name,last_name,email,password,status) values (?,?,?,?,?,?)"
	result, err := r.db.Exec(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Password, user.Status)

	if err != nil {
		return 0, err
	}

	lastInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(lastInserted), nil

}

func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {

	query := "UPDATE users SET "
	var args []any
	var sets []string

	if user.UserName != "" {
		sets = append(sets, "username = ?")
		args = append(args, user.UserName)
	}
	if user.FirstName != "" {
		sets = append(sets, "first_name = ?")
		args = append(args, user.FirstName)
	}
	if user.LastName != "" {
		sets = append(sets, "last_name = ?")
		args = append(args, user.LastName)
	}
	if user.Email != "" {
		sets = append(sets, "email = ?")
		args = append(args, user.Email)
	}
	if user.Password != "" {
		sets = append(sets, "password = ?")
		args = append(args, user.Password)
	}
	if user.Status != "" {
		sets = append(sets, "status = ?")
		args = append(args, user.Status)
	}

	if len(sets) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(sets, ", ") + " WHERE id = ?"
	args = append(args, user.ID)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return r.GetUserByUsername(user.UserName)
}

func (r *UserRepository) ExcistsByUsernameAndEmail(username string, email string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = ? OR email = ?)"
	var exist bool
	err := r.db.QueryRow(query, username, email).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil

}
