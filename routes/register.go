package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/books", func(ctx *gin.Context) {
		fmt.Println("you got all books")
	})

	app.GET("/books/:id", func(ctx *gin.Context) {
		fmt.Println("you got a book")
	})

	app.POST("/books", func(ctx *gin.Context) {
		fmt.Println("you created a book")
	})

	app.PUT("/books/:id", func(ctx *gin.Context) {
		fmt.Println("you edited a book")
	})

	app.DELETE("/books/:id", func(ctx *gin.Context) {
		fmt.Println("you deleted a book")
	})
}
