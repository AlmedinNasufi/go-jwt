package initializers

import "github.com/AlmedinNasufi/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
