package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/pkg/pg"
)

type BasketPostgres struct {
	db *sql.DB
}

func NewBasketPostgres(db *sql.DB) *BasketPostgres {
	return &BasketPostgres{db: db}
}

func (p *BasketPostgres) GetBasket(ctx context.Context, userId int) (models.Basket, error) {
	var basket models.Basket

	query := fmt.Sprintf("SELECT book_id FROM %s WHERE user_id=$1", pg.BasketTable)
	rows, err := p.db.Query(query, userId)
	if err != nil {
		return basket, err
	}

	for rows.Next() {
		var bookId int
		var book models.Book

		if err = rows.Scan(&bookId); err != nil {
			return basket, err
		}

		query = fmt.Sprintf("SELECT * FROM %s WHERE id=$1", pg.BooksTable)

		row, err1 := p.db.Query(query, bookId)
		if err1 != nil {
			return basket, err1
		}

		for row.Next() {
			if err = row.Scan(
				&book.Id,
				&book.Title,
				&book.YearPublication,
				&book.Author,
				&book.Price,
				&book.QtyStock,
				&book.CategoryId,
			); err != nil {
				return basket, err
			}

		}

		if book.QtyStock != 0 && book.CategoryId != 0 {
			basket.Books = append(basket.Books, book)
		}
	}

	return basket, err
}

func (p *BasketPostgres) ToPayBasket(ctx context.Context, userId int) error {
	query := fmt.Sprintf(
		"SELECT book_id FROM %s WHERE user_id=$1", pg.BasketTable,
	)

	rows, err := p.db.Query(query, userId)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(
		"DELETE FROM %s WHERE user_id=$1", pg.BasketTable,
	)
	if _, err = p.db.Exec(query, userId); err != nil {
		return err
	}

	for rows.Next() {
		var bookId int

		if err = rows.Scan(&bookId); err != nil {
			return err
		}

		query = fmt.Sprintf(
			"UPDATE %s SET qty_stock=qty_stock-1 WHERE id=$1", pg.BooksTable,
		)
		if _, err = p.db.Exec(query, bookId); err != nil {
			return err
		}
	}

	return nil
}

func (p *BasketPostgres) UpdateBasket(ctx context.Context, userId, bookId int, method string) (models.Basket, error) {
	if method == "DELETE" {
		query := fmt.Sprintf(
			"DELETE FROM %s WHERE user_id=$1 AND book_id=$2", pg.BasketTable,
		)

		if _, err := p.db.Exec(query, userId, bookId); err != nil {
			return models.Basket{}, err
		}

	} else if method == "PUT" {
		var basket models.Basket

		query := fmt.Sprintf(
			"SELECT * FROM %s WHERE user_id=$1 AND book_id=$2", pg.BasketTable,
		)

		rows, err := p.db.Query(query, userId, bookId)
		if err != nil {
			return basket, err
		}
		for rows.Next() {
			if err = rows.Scan(&basket.UserId, &basket.BookId); err != nil {
				return basket, err
			}
		}

		query = fmt.Sprintf(
			"INSERT INTO %s (user_id, book_id) VALUES ($1, $2)", pg.BasketTable,
		)
		if _, err := p.db.Exec(query, userId, bookId); err != nil {
			return basket, errors.New("book already in basket")
		}

	}

	return p.GetBasket(ctx, userId)
}
