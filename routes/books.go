package routes

import (
	"fmt"
	"goreads/models"
	"goreads/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllBooks(ctx *gin.Context) {
	books, err := models.GetAllBooks()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": "Cannot serve your request",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		"message": "All books in the database",
		"items":   books,
	})
}

func getOneBook(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": "Cannot read path variable :id",
		})
		return
	}

	book, err := models.GetOneBook(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": fmt.Sprintf("Cannot get book with id %d", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		"message": fmt.Sprintf("Book with id %d", id),
		"item":    &book,
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
