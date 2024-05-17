package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/pkg/pg"
)

type Category interface {
	GetCategories(ctx context.Context) ([]models.Category, error)
	GetCategory(ctx context.Context, catId int) (models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) (models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category) (models.Category, error)
	DeleteCategory(ctx context.Context, catId int) error
}

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
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
	err := r.db.Get(&category, query, catId)

	return category, err
}

func (r *CategoryPostgres) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	// сначала проверяю, есть ли в базе уже эта категория
	if _, ok := r.checkCatInTable(ctx, category.Title); ok == nil {
		return models.Category{}, errors.New("category already exists")
	}

	// если нет, то заполняю таблицу новой категорией
	query := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1)", pg.CategoriesTable)
	if _, err := r.db.Exec(query, category.Title); err != nil {
		return models.Category{}, err
	}

	// после заполнения отдаю эту категорию для вывода
	return r.checkCatInTable(ctx, category.Title)
}

func (r *CategoryPostgres) UpdateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	// сначала проверяю, есть ли в базе уже эта категория на которую я хочу изменить текущую
	if _, ok := r.checkCatInTable(ctx, category.Title); ok == nil {
		return category, errors.New("category already exists")
	}

	// если нет, то обновляю в таблице категорию
	query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2", pg.CategoriesTable)
	if _, err := r.db.Exec(query, category.Title, category.Id); err != nil {
		return models.Category{}, err
	}

	// после обновления отдаю эту категорию для вывода
	return r.checkCatInTable(ctx, category.Title)
}

func (r *CategoryPostgres) checkCatInTable(ctx context.Context, title string) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf("SELECT * FROM %s WHERE title=$1", pg.CategoriesTable)
	ok := r.db.Get(&category, query, title)

	return category, ok
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
