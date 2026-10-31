package pgsql

import (
	"database/sql"
	"errors"

	"clean-go/internal/entity"
	"clean-go/internal/repository"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepoPG(db *sqlx.DB) repository.UserRepository {
	return &userRepo{db: db}
}

// Get all users
func (r *userRepo) All() ([]*entity.User, error) {
	var users []*entity.User

	err := r.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Get user by id
func (r *userRepo) GetById(id string) (*entity.User, error) {
	var user entity.User

	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Get user by email
func (r *userRepo) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create user
func (r *userRepo) Create(user *entity.User) error {
	query := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())"
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	return err
}
