package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/pkg/pg"
)

type User interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUser(ctx context.Context, email, password string) (models.User, error)
	IsAdmin(ctx context.Context, userId int) (models.User, error)
}

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, newUser models.User) (int, error) {
	var user models.User
	// проверяю зарегистрирован ли данный email ранее
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1", pg.UsersTable)

	// если такой email уже есть в бд, то возвращаю ошибку
	if ok := r.db.Get(&user, query, newUser.Email); ok == nil {
		return 0, errors.New("email has already been registered")
	}

	// если пользователя под таким email нет в базе, то заполняю таблицу новыми полями
	query = fmt.Sprintf(
		"INSERT INTO %s (username, email, password) values ($1, $2, $3) RETURNING id",
		pg.UsersTable,
	)
	row := r.db.QueryRow(query, newUser.Username, newUser.Email, newUser.Password)

	// получаю из запроса id нового пользователя и возвращаю его как результат функции
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUser(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", pg.UsersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *UserPostgres) IsAdmin(ctx context.Context, userId int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT is_admin FROM %s WHERE id=$1", pg.UsersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}
