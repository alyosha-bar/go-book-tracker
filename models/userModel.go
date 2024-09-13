package models

type User struct {
	User_id  int
	Username string
	Email    string `gorm:"unique"`
	Password string
}
