package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/pkg/pg"
)

type Book interface {
	GetBooks(ctx context.Context) ([]models.Book, error)
	GetBook(ctx context.Context, bookId int) (models.Book, error)
}

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) GetBooks(ctx context.Context) ([]models.Book, error) {
	var book models.Book
	var books []models.Book

	query := fmt.Sprintf("SELECT * FROM %s", pg.BooksTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(
			&book.Id,
			&book.Title,
			&book.YearPublication,
			&book.Author,
			&book.Price,
			&book.QtyStock,
			&book.CategoryId,
		); err != nil {
			return nil, err
		}
		if book.QtyStock != 0 {
			books = append(books, book)
		}

	}

	return books, nil
}

func (r *BookPostgres) GetBook(ctx context.Context, bookId int) (models.Book, error) {
	var book models.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", pg.BooksTable)
	err := r.db.Get(&book, query, bookId)
	if book.QtyStock == 0 {
		return models.Book{}, errors.New("book not found")
	}

	return book, err
}
