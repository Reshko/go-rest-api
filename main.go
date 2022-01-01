package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBook)
	router.POST("/books/:id", getBookById)
	router.DELETE("/books/:id", deleteBookById)

	router.Run("localhost:8080")

}

type book struct {
	ID     string  `json"id"`
	Title  string  `json"title"`
	Writer string  `json"writer"`
	Price  float64 `json"float64"`
}

var books = []book{
	{ID: "1", Title: "Blue Train", Writer: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Writer: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Writer: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Blue Train", Writer: "John Coltrane", Price: 56.99},
	{ID: "5", Title: "Jeru", Writer: "Gerry Mulligan", Price: 17.99},
	{ID: "6", Title: "Sarah Vaughan and Clifford Brown", Writer: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func getBookById(c *gin.Context) {
	var id = c.Param("ID")

	for _, element := range books {
		if element.ID == id {
			c.IndentedJSON(http.StatusOK, element)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "non found"})
}

func deleteBookById(c *gin.Context) {
	var id = c.Param("id")

	for index, value := range books {
		if value.ID == id {
			books = append(books[:index], books[:index+1]...)
			c.IndentedJSON(http.StatusOK, books)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "non found"})
}
