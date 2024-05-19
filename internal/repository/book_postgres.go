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
	CreateBook(ctx context.Context, book models.Book) (models.Book, error)
	UpdateBook(ctx context.Context, book models.Book) (models.Book, error)
	DeleteBook(ctx context.Context, bookId int) error
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
		if book.QtyStock != 0 && book.CategoryId != 0 {
			books = append(books, book)
		}

	}

	return books, nil
}

func (r *BookPostgres) GetBook(ctx context.Context, bookId int) (models.Book, error) {
	var book models.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", pg.BooksTable)
	err := r.db.Get(&book, query, bookId)
	fmt.Println(book)
	if book.QtyStock == 0 || book.CategoryId == 0 {
		return models.Book{}, errors.New("book not found")
	}

	return book, err
}

func (r *BookPostgres) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {
	// сначала проверяю, есть ли в базе уже эта книга
	if _, ok := r.checkBookInTable(ctx, book.Title); ok == nil {
		return models.Book{}, errors.New("book already exists")
	}

	// если нет, то заполняю таблицу новой книгой
	query := fmt.Sprintf(
		`INSERT INTO %s (title, year_publication, author, price, qty_stock, category_id) 
				VALUES ($1, $2, $3, $4, $5, $6)`,
		pg.BooksTable,
	)
	if _, err := r.db.Exec(
		query, book.Title, book.YearPublication, book.Author, book.Price, book.QtyStock, book.CategoryId,
	); err != nil {
		return models.Book{}, err
	}

	// после заполнения отдаю эту книгу для вывода
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

	// после обновления отдаю эту книгу для вывода
	return r.checkBookInTable(ctx, book.Title)
}

func (r *BookPostgres) checkBookInTable(ctx context.Context, title string) (models.Book, error) {
	var book models.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE title=$1", pg.BooksTable)
	ok := r.db.Get(&book, query, title)

	return book, ok
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
