package main

import (
	"errors"
	//"errors"
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
	{"1", "Half Girlfriend", "Chetan Bhagat", 1},
	{"2", "One Indian Girl", "Chetan Bhagat", 2},
	{"3", "Sharp Objects", "Gillian Flynn", 1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(id string, c *gin.Context) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id, c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,"Id not found")
		panic(err)
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest,"Id not found")
		panic(err)
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	id:= c.Param("id")
	err := delBookById(id, c)
	if err != nil {
		panic(err)
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func delBookById(id string, c *gin.Context) error {
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.DELETE("/books/delete/:id", deleteBook)
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
		return
	}
}


