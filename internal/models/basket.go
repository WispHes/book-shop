package models

type Basket struct {
	Id      int    `json:"id"`
	UserId  string `json:"user_id"`
	BookId  string `json:"book_id" binding:"required"`
	DateAdd string `json:"date_add"`
}
