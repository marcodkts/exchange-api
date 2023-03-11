package initializers

import "exchange-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Log{})
}