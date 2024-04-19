package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{"1","Half Girlfriend","Chetan Bhagat",1},
	{"2","One Indian Girl","Chetan Bhagat",2},
	{"3","Sharp Objects","Gillian Flynn",1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router:=gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	err := router.Run("localhost:8080")
	if err != nil {
		return 
	}
}
