package models

type User_Book struct {
	id      int
	user_id string
	book_id string `gorm:"unique"`
	status  int
}
