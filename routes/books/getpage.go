package books

import (
	"fmt"
	"goreads/models"
	"goreads/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooksPage(ctx *gin.Context) {
	pageParam, isPresent := ctx.GetQuery("page")

	var page int

	if !isPresent {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pageParam)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.Response{
				"message": "Inavalid value for page query param",
			})
			return
		}
	}

	if page <= 0 {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": "Page number should be a positive integer",
		})
		return
	}

	books, totalRows, selectedRows, totalPages, err := models.GetBooksPage(page)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			"message": "Cannot serve your request",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		"items":       books,
		"totalItems":  totalRows,
		"itemsInPage": selectedRows,
		"page":        page,
		"totalPages":  totalPages,
	})
}
