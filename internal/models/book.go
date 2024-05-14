package models

type Book struct {
	Id              int    `json:"id" db:"id"`
	Title           string `json:"title" binding:"required"`
	YearPublication string `json:"year" binding:"required"`
	Author          string `json:"author" binding:"required"`
	Price           string `json:"price" binding:"required"`
	QtyStock        string `json:"qtyStock" binding:"required"`
	CategoryId      string `json:"category_id" binding:"required"`
}

type Category struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" binding:"required" db:"title"`
}
