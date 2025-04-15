package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/matteusmoreno/gopportunities/schemas"
	"net/http"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID is required",
		})
		return
	}

	opening := schemas.Opening{}

	err := db.First(&opening, id).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Opening not found",
		})
		return
	}

	err = db.Delete(&opening).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
}
