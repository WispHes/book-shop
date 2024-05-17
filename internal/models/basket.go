package models

type Basket struct {
	UserId int    `json:"user_id" db:"user_id"`
	BookId int    `json:"book_id" binding:"required" db:"book_id"`
	Books  []Book `json:"books"`
}
