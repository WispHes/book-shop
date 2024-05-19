package models

type Book struct {
	Id              int     `json:"id" db:"id"`
	Title           string  `json:"title" binding:"required" db:"title"`
	YearPublication int     `json:"year_publication" binding:"required" db:"year_publication"`
	Author          string  `json:"author" binding:"required" db:"author"`
	Price           float64 `json:"price" binding:"required" db:"price"`
	QtyStock        int     `json:"qty_stock" binding:"required" db:"qty_stock"`
	CategoryId      int     `json:"category_id" binding:"required" db:"category_id"`
}

type Category struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" binding:"required" db:"title"`
}
