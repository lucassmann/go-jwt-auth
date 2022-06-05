package database

import (
	"go_auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:rootroot@/go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connector to the database")
	}

	connection.AutoMigrate(&models.User{})
}
