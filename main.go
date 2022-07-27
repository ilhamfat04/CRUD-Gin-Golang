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
	fmt.Println("Databace connect")

	// ================== RAW QUERY ====================
	// ################## INSERT ##################
	// err = db.Exec("INSERT INTO books(title,description,price,created_at,updated_at) VALUES ('Manusia Hariamu','Test',2000,?,?)", time.Now(), time.Now()).Error
	// ############################################

	// ############# SELECT SINGLE ##############
	// var book book.Book
	// db.Raw("SELECT * FROM books WHERE id=?", 4).Scan(&book)
	// fmt.Printf("Data : %v", book)
	// ############################################

	// ############# SELECT MULTIPLE ##############
	// var books []book.Book
	// db.Raw("SELECT * FROM books").Scan(&books)

	// for _, b := range books {
	// 	fmt.Printf("Data : %v", b)
	// }
	// ############################################

	// ############# UPDATE DATA ##############
	// var book []book.Book
	// db.Raw("UPDATE books SET title=? WHERE id=?", "Man of Body Tiger", 3).Scan(&book)
	// ########################################

	// ############# DELETE DATA ##############
	// db.Exec("DELETE FROM books WHERE id=?", 6)
	// ########################################
	// =================================================

	// ================== ORM QUERY ====================
	// ################### INSERT ###################
	// book := book.Book{
	// 	Title:       "Manusia Hariamu",
	// 	Description: "Buku tentang seorang manusia yang menjadi harimau",
	// 	Price:       200000,
	// }
	// err = db.Create(&book).Error
	// ###############################################

	// ############# SELECT SINGLE DATA #############
	// var book book.Book // store single data
	// err = db.Take(&book, 4).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("|| Error creating book record ||")
	// 	fmt.Println("================================")
	// }
	// fmt.Println("Title : ", book.Title) //single data
	// fmt.Printf("Data : %v", book)       //single data
	// ###############################################

	// ############ SELECT MULTIPLE DATA #############
	// var books []book.Book // store multiple data
	// err = db.Find(&books).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("|| Error creating book record ||")
	// 	fmt.Println("================================")
	// }

	// for _, b := range books {
	// 	fmt.Printf("Data : %v", b) //multiple data
	// }
	// ###############################################

	// ################ UPDATE DATA ##################
	// var book book.Book

	// err = db.Take(&book, 3).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("|| Error creating book record ||")
	// 	fmt.Println("================================")
	// }

	// book.Title = "Man Tiger Wong"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("|| Error creating book record ||")
	// 	fmt.Println("================================")
	// }
	// ###############################################

	// ################ DELETE DATA ##################
	// var book book.Book
	// db.Take(&book, 3)

	// db.Delete(&book)
	// ###############################################

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.AddBookHandler)

	router.Run("127.0.0.1:8000")
}
