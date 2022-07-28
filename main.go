package main

import (
	"golang-api/book"
	"golang-api/handler"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}

	db.AutoMigrate(&book.Book{}) // migration

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler((bookService))

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/book/:id", bookHandler.GetBook)
	v1.PATCH("/book/:id", bookHandler.UpdateBook)

	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.AddBookHandler)

	router.Run("127.0.0.1:8000")
}
