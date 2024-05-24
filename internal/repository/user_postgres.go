package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/pkg/pg"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, newUser models.User) (int, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (username, email, password) values ($1, $2, $3) RETURNING id",
		pg.UsersTable,
	)
	row := r.db.QueryRow(query, newUser.Username, newUser.Email, newUser.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUser(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", pg.UsersTable)
	row := r.db.QueryRow(query, email, password)
	err := row.Scan(&user.Id)

	return user, err
}

func (r *UserPostgres) IsAdmin(ctx context.Context, userId int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT is_admin FROM %s WHERE id=$1", pg.UsersTable)
	rows, err := r.db.Query(query, userId)
	for rows.Next() {
		err = rows.Scan(&user.IsAdmin)
	}

	return user, err
}
