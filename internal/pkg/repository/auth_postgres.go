package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/wisphes/book-shop/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

const usersTable = "users"

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(ctx context.Context, newUser models.User) (int, error) {
	// проверяю зарегистрирован ли данный email ранее
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1", usersTable)
	err := r.db.Get(&user, query, newUser.Email)

	// если такой email уже есть в бд, то возвращаю ошибку
	if err == nil {
		return 0, errors.New("email has already been registered")
	}

	// если пользователя под таким email нет в базе, то заполняю таблицу новыми полями
	query = fmt.Sprintf(
		"INSERT INTO %s (username, email, password) values ($1, $2, $3) RETURNING id",
		usersTable,
	)
	row := r.db.QueryRow(query, newUser.Username, newUser.Email, newUser.Password)

	// получаю из запроса id нового пользователя и возвращаю его как результат функции
	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *AuthPostgres) IsAdmin(ctx context.Context, id int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT is_admin FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}
