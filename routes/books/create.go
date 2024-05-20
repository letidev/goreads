package books

import (
	"goreads/models"
	"goreads/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var bookObj models.Book
	err := ctx.ShouldBindJSON(&bookObj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": "Invalid object",
		})
		return
	}

	err = bookObj.SaveNew()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
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
