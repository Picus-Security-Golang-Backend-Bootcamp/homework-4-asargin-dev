package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/factory"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/models"
	"github.com/julienschmidt/httprouter"
)

var books []models.Book
var repository *factory.Repository

func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w,
		`Kitapp API v1'e Hoşgeldiniz

		Bu uygulamadan faydalanabileceğiniz servisler:

		/list - Kütüphanedeki tüm kitapları gösterir
		/search?{query} - Aramak istediğiniz kitapları gösterir
		/buy?book_id={id}?quantity{number} - Kitap satın almayı sağlar
		/delete?id={id} - Kütüphaneden kitap silme işlemini sağlar.  `)
}

func ListBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	books = repository.FindAll()

	result, _ := json.Marshal(books)    //get json byte array
	length := len(result)               //Find the length of the byte array
	response := string(result[:length]) //convert to string
	fmt.Fprint(w, response)             //write to response
}

func SearchBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	query := r.URL.Query()
	id := query.Get("id")
	query_id, _ := strconv.Atoi(id)
	book := repository.FindById(query_id)

	result, _ := json.Marshal(book)
	length := len(result)
	response := string(result[:length])
	fmt.Fprint(w, response)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	query := r.URL.Query()
	id := query.Get("id")
	query_id, _ := strconv.Atoi(id)
	book := repository.DeleteById(query_id)

	result, _ := json.Marshal(book)
	length := len(result)
	response := string(result[:length])
	fmt.Fprint(w, response)
}
