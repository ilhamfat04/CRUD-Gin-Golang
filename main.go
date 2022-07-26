package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/books/:id/:title", booksHandler)
	v1.GET("/query", queryHandler)

	v1.POST("/books", addBookHandler)

	router.Run("127.0.0.1:8000")
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":    "ilham fathullah",
		"address": "Jambi",
	})
}

// URL PARAMS
func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

// QUERY STRING
func queryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{"title": title})
}

// struct for store data post
// no casesensitive for  variable name when submit form JSON postman
// 		subTitle == SubTitle, etc
type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `binding:"required,number"`
	// SubTitle string `json:"sub_title"`  // when different var name and json form
	SubTitle string
}

func addBookHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, conditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":    bookInput.Title,
		"price":    bookInput.Price,
		"subTitle": bookInput.SubTitle,
	})
}
