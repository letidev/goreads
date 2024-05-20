package routes

import (
	"goreads/models"
	"goreads/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllBooks(ctx *gin.Context) {
	books, err := models.GetAllBooks()

	if err != nil {
		ctx.JSON(400, utils.Response{
			"message": "Cannot serve your request",
		})
		return
	}

	ctx.JSON(200, utils.Response{
		"message": "All books in the database",
		"items":   books,
	})
}

func getOneBook(ctx *gin.Context) {
	ctx.JSON(200, utils.Response{
		"message": "you got one book",
	})
}

func createBook(ctx *gin.Context) {
	var bookObj models.Book
	err := ctx.ShouldBindJSON(&bookObj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": "Invalid object",
		})
		return
	}

	err = bookObj.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": "Could not insert book object",
			"err":     err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response{
		"message": "Book created successfully",
		"item":    bookObj,
	})
}

func editBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, utils.Response{
		"message": "you edited a book",
	})
}

func deleteBook(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, utils.Response{
		"message": "you deleted a book",
	})
}
