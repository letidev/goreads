package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/books", getAllBooks)

	app.GET("/books/:id", getOneBook)

	app.POST("/books", createBook)

	app.PUT("/books/:id", editBook)

	app.DELETE("/books/:id", deleteBook)
}
