package models

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"todo/password"
	"todo/storage"
)

// User model
type User struct {
	ID        string `json:"id" form:"id"`
	Firstname string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
}

// UserFilter items for filtering task
type UserFilter struct {
	Email    string
	Password string
}

// UserRepository operations for persistance layer
type UserRepository struct {
	db      *sqlx.DB
	storage storage.Adapter
}

// Login get user given email and plain password
func (repository *UserRepository) Login(email, pwd string) (*User, error) {
	user := &User{}
	sql := "SELECT * FROM users WHERE email = $1"

	if err := repository.db.Get(user, sql, email); err != nil {
		return nil, errors.New("User not found")
	}

	if !password.ComparePasswords(user.Password, []byte(pwd)) {
		return nil, errors.New("Invalid password")
	}

	return user, nil
}

// Create insert a new user
func (repository *UserRepository) Create(data *User) (*User, error) {
	hashedPassword, err := password.HashAndSalt([]byte(data.Password))
	if err != nil {
		return nil, err
	}

	data.Password = hashedPassword

	statement, err := repository.db.PrepareNamed(`
		INSERT INTO users (email, firstname, lastname, password)
		VALUES (:email, :firstname, :lastname, :password)
		RETURNING id
	`)

	if err != nil {
		return nil, err
	}

	err = statement.Get(&data.ID, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewUserRepository exposes User repository methods
func NewUserRepository(db *sqlx.DB, storage storage.Adapter) *UserRepository {
	return &UserRepository{db: db, storage: storage}
}
