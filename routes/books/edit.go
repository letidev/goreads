package books

import (
	"goreads/models"
	"goreads/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditBook(ctx *gin.Context) {
	var bookObj models.Book
	err := ctx.ShouldBindJSON(&bookObj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": "Invalid object",
		})
		return
	}

	err = bookObj.SaveExisting()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			"message": "Could not save item",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		"message": "Item saved successfully",
		"item":    bookObj,
	})
}
