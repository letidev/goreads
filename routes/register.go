package routes

import (
	"goreads/routes/books"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/books", books.GetBooksPage)

	app.GET("/books/:id", books.GetOneBook)

	app.POST("/books", books.CreateBook)

	app.PUT("/books", books.EditBook)

	app.DELETE("/books/:id", books.DeleteBook)
}
