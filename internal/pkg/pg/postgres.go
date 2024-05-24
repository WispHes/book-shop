package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	CategoriesTable = "categories"
	UsersTable      = "users"
	BooksTable      = "books"
	BasketTable     = "basket"
)

type Config struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	//m, err := migrate.NewWithDatabaseInstance(
	//	"file:///migrations",
	//	"postgres", driver)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
