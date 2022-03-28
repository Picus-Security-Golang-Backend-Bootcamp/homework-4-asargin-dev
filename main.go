package main

import (
	"net/http"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/factory"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/router"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/utils"
)

var bookList []models.Book
var repository *factory.Repository

func init() {
	db := factory.ConnectDatabase("postgres://asargin:m@localhost:5432/postgis")
	repository = factory.NewRepository(db)
	repository.Migrate()

	bookList, _ = utils.CsvReader("docs/books.csv")
	repository.InsertData(bookList)
}

func main() {
	http.ListenAndServe(":8088", router.Route())
}
