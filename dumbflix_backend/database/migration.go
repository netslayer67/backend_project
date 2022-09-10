package database

import (
	"fmt"
	"testdumpflix/models"
	"testdumpflix/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Profile{},
		&models.User{},
		&models.Category{},
		&models.Film{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
