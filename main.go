package main

import (
	"errors" // Add this import for the errors package
	"github.com/gin-gonic/gin"
	"net/http"
)

// Define Book struct if not already defined
type Book struct {
	Title  string
	Author string
	// Add other fields as needed
}

var library []Book

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*") // Assuming your templates are in a "templates" folder

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"library": library})
	})

	router.POST("/add", func(c *gin.Context) {
		book, err := createBook(c.PostForm("title"), c.PostForm("author"))
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": err.Error()})
			return
		}
		library = append(library, book)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}

func createBook(title, author string) (Book, error) {
	// You can add validation logic here if needed
	if title == "" || author == "" {
		return Book{}, errors.New("Title and Author are required")
	}

	book := Book{
		Title:  title,
		Author: author,
		// Initialize other fields as needed
	}

	return book, nil
}
