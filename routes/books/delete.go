package books

import (
	"fmt"
	"goreads/models"
	"goreads/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			"message": "Cannot read path variable :id",
		})
		return
	}

	err = models.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			"message": fmt.Sprintf("Could not delete record with id %d", id),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, utils.Response{})
}
