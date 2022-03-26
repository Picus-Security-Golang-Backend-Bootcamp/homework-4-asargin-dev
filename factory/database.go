package factory

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(conString string) *gorm.DB {
	dns := "postgres://asargin:@localhost:5432/postgis"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database : %s", err.Error()))
	}

	db.AutoMigrate(&models.Book{})

	if err != nil {
		panic(err)
	}

	return db
}
