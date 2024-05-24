package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/pkg/pg"
)

type BookPostgres struct {
	db *sql.DB
}

func NewBookPostgres(db *sql.DB) *BookPostgres {
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
		if book.QtyStock != 0 && book.CategoryId != 0 {
			books = append(books, book)
		}

	}

	return books, nil
}

func (r *BookPostgres) GetBook(ctx context.Context, bookId int) (models.Book, error) {
	var book models.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", pg.BooksTable)
	row := r.db.QueryRow(query, bookId)
	if err := row.Scan(
		&book.Id,
		&book.Title,
		&book.YearPublication,
		&book.Author,
		&book.Price,
		&book.QtyStock,
		&book.CategoryId,
	); err != nil {
		return book, err
	}

	if book.QtyStock == 0 || book.CategoryId == 0 {
		return models.Book{}, errors.New("book not found")
	}

	return book, nil
}

func (r *BookPostgres) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {
	query := fmt.Sprintf(
		`INSERT INTO %s (title, year_publication, author, price, qty_stock, category_id)
				VALUES ($1, $2, $3, $4, $5, $6)`,
		pg.BooksTable,
	)
	if _, err := r.db.Exec(
		query, book.Title, book.YearPublication, book.Author, book.Price, book.QtyStock, book.CategoryId,
	); err != nil {
		return models.Book{}, errors.New("book already exists")
	}

	return r.checkBookInTable(ctx, book.Title)
}

func (r *BookPostgres) UpdateBook(ctx context.Context, book models.Book) (models.Book, error) {
	query := fmt.Sprintf(
		`UPDATE %s
			    SET title=$1, year_publication=$2, author=$3, price=$4, qty_stock=$5, category_id=$6
			    WHERE id=$7`,
		pg.BooksTable,
	)
	if _, err := r.db.Exec(
		query, book.Title, book.YearPublication, book.Author, book.Price, book.QtyStock, book.CategoryId, book.Id,
	); err != nil {
		return models.Book{}, err
	}

	return r.checkBookInTable(ctx, book.Title)
}

func (r *BookPostgres) checkBookInTable(ctx context.Context, title string) (models.Book, error) {
	var book models.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE title=$1", pg.BooksTable)

	row := r.db.QueryRow(query, title)
	if err := row.Scan(
		&book.Id,
		&book.Title,
		&book.YearPublication,
		&book.Author,
		&book.Price,
		&book.QtyStock,
		&book.CategoryId,
	); err != nil {
		return book, err
	}

	return book, nil
}

func (r *BookPostgres) DeleteBook(ctx context.Context, bookId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", pg.BooksTable)
	if _, err := r.db.Exec(query, bookId); err != nil {
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE book_id=$1", pg.BasketTable)
	if _, err := r.db.Exec(query, bookId); err != nil {
		return err
	}
	return nil
}
