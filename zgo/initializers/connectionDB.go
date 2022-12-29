package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	var err error
	dsn := "host=lucky.db.elephantsql.com user=yifqwxhu password=sAh2My3Mk_yKR4fx6HS4kBloLrvzUaL9 dbname=yifqwxhu port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
