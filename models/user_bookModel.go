package models

type User_Book struct {
	ID      int
	User_id int
	Book_id int `gorm:"unique"`
	Status  int
}
