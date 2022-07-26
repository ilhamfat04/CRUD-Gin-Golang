package main

import (
	"fmt"
	"golang-api/book"
	"golang-api/handler"
	"log"
	"time"

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
	fmt.Println("Databace connect")

	// book := book.Book{
	// 	Title:       "Manusia Hariamu",
	// 	Description: "Buku tentang seorang manusia yang menjadi harimau",
	// 	Price:       200000,
	// }

	err = db.Exec("INSERT INTO books(title,description,price,created_at,updated_at) VALUES ('Manusia Hariamu','Test',2000,?,?)", time.Now(), time.Now()).Error

	// err = db.Create(&book).Error
	if err != nil {
		fmt.Println("===============================")
		fmt.Println("|| Error creating boo record ||")
		fmt.Println("===============================")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.AddBookHandler)

	router.Run("127.0.0.1:8000")
}
