package repo

import (
	"database/sql"
	"fmt"
)

type AuthRepositoryInterface interface {
	GetHashedPassword(username string) (string, error)
}

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepositoryInterface {
	return &AuthRepo{db: db}
}

func (ar *AuthRepo) GetHashedPassword(username string) (string, error) {
	query := "SELECT password from users where username = ?"
	var password string
	err := ar.db.QueryRow(query, username).Scan(&password)
	if err != nil {
		return "", fmt.Errorf("user/password not found")
	}
	return password, nil
}
