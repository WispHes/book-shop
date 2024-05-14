package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/wisphes/book-shop/internal/models"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

const categoriesTable = "categories"

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) GetCategories(ctx context.Context) ([]models.Category, error) {
	var category models.Category
	var categories []models.Category

	query := fmt.Sprintf("SELECT * FROM %s", categoriesTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Title)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryPostgres) GetCategory(ctx context.Context, catId int) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", categoriesTable)
	err := r.db.Get(&category, query, catId)

	return category, err
}

func (r *CategoryPostgres) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	// сначала проверяю, есть ли в базе уже эта категория
	_, err := r.checkCatInTable(ctx, 1, category.Title)
	if err != nil {
		return models.Category{}, err
	}

	// если нет, то заполняю таблицу новой категорией
	query := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1)", categoriesTable)
	_, err = r.db.Exec(query, category.Title)
	if err != nil {
		return models.Category{}, err
	}

	// после заполнения отдаю эту категорию для вывода
	return r.checkCatInTable(ctx, 2, category.Title)
}

func (r *CategoryPostgres) UpdateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	// сначала проверяю, есть ли в базе уже эта категория на которую я хочу изменить текущую
	_, err := r.checkCatInTable(ctx, 1, category.Title)
	if err != nil {
		return category, err
	}

	// если нет, то обновляю в таблице категорию
	query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2", categoriesTable)
	_, err = r.db.Exec(query, category.Title, category.Id)
	if err != nil {
		return models.Category{}, err
	}

	// после обновления отдаю эту категорию для вывода
	return r.checkCatInTable(ctx, 2, category.Title)
}

// Функция обрабатывает 2 случая:
// Случай - параметр в функции flag 1/2
// 1 - существует ли категория в бд которую мы хотим создать/изменить
// 2 - получить измененную/созданную категорию
func (r *CategoryPostgres) checkCatInTable(ctx context.Context, flag int8, title string) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf("SELECT * FROM %s WHERE title=$1", categoriesTable)
	err := r.db.Get(&category, query, title)

	// обработка первого случая
	if flag == 1 && err == nil {
		return category, errors.New("category already exists")
	}

	return category, nil
}
