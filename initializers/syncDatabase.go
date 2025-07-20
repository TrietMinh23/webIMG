package initializers

import "webimg/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
