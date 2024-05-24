package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/pkg/pg"
)

type CategoryPostgres struct {
	db *sql.DB
}

func NewCategoryPostgres(db *sql.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) GetCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	query := fmt.Sprintf("SELECT * FROM %s", pg.CategoriesTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category models.Category

		if err = rows.Scan(&category.Id, &category.Title); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryPostgres) GetCategory(ctx context.Context, catId int) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", pg.CategoriesTable)

	rows := r.db.QueryRow(query, catId)
	if err := rows.Scan(&category.Id, &category.Title); err != nil {
		return category, errors.New("category is not in table")
	}

	return category, nil
}

func (r *CategoryPostgres) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	query := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1)", pg.CategoriesTable)
	if _, err := r.db.Exec(query, category.Title); err != nil {
		return models.Category{}, errors.New("category already exists")
	}

	return r.checkCatInTable(ctx, category.Title)
}

func (r *CategoryPostgres) UpdateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2", pg.CategoriesTable)
	if _, err := r.db.Exec(query, category.Title, category.Id); err != nil {
		return models.Category{}, err
	}

	return r.checkCatInTable(ctx, category.Title)
}

func (r *CategoryPostgres) checkCatInTable(ctx context.Context, title string) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf("SELECT * FROM %s WHERE title=$1", pg.CategoriesTable)

	rows := r.db.QueryRow(query, title)
	if err := rows.Scan(&category.Id, &category.Title); err != nil {
		return category, err
	}

	return category, nil
}

func (r *CategoryPostgres) DeleteCategory(ctx context.Context, catId int) error {
	query := fmt.Sprintf("UPDATE %s SET category_id=0 WHERE category_id=$1", pg.BooksTable)
	if _, err := r.db.Exec(query, catId); err != nil {
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", pg.CategoriesTable)
	if _, err := r.db.Exec(query, catId); err != nil {
		return err
	}

	return nil
}
