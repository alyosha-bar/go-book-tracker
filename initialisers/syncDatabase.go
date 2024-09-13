package initialisers

import "books-backend/models"

func SyncDatabase() {

	DB := ConnectToDB()

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User_Book{})
}
