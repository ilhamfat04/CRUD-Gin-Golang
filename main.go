package main

import (
	"fmt"
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

	// ############## ADD BOOK ###############
	// bookInput := book.BookInput{
	// 	Title:       "Who Came After Die?",
	// 	Description: "Awesome book",
	// 	Price:       90000,
	// }

	// bookService.AddBook(bookInput)

	// fmt.Println("Title : ", bookInput.Title)
	// #######################################

	// ############## GET BOOKS ###############
	var books []book.Book
	books, _ = bookService.GetBooks()
	fmt.Println("Title : ", books)

	for _, book := range books {
		fmt.Println("Title : ", book.Title)
	}
	// ########################################

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.AddBookHandler)

	router.Run("127.0.0.1:8000")
}
