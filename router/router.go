package router

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/handler"
	"github.com/julienschmidt/httprouter"
)

func Route() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", handler.Index)
	router.GET("/list", handler.ListBooks)
	router.GET("/search", handler.SearchBooks)

	return router
}
