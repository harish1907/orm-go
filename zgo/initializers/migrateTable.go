package initializers

import "github.com/harish1907/zgo/models"

func MigrationTable() {
	DB.AutoMigrate(
		&models.State{},
		&models.Village{},
		&models.MyUser{},
		&models.Language{},
	)
}
