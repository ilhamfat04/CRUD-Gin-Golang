package main

import (
	"golang-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.AddBookHandler)

	router.Run("127.0.0.1:8000")
}