package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Book1", Author: "A1", Quantity: 10},
	{ID: "2", Title: "Book2", Author: "A2", Quantity: 20},
	{ID: "3", Title: "Book3", Author: "A3", Quantity: 30},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByIdHelper(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByIdHelper(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBookQuantity(c *gin.Context) {
	id := c.Param("id")
	quantity := c.Param("quantity")
	book, err := getBookByIdHelper(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	book.Quantity, _ = strconv.Atoi(quantity)
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.PATCH("/updateBooks/:id/:quantity", updateBookQuantity)
	router.Run("localhost:8080")
}
