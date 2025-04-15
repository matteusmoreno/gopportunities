package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/matteusmoreno/gopportunities/schemas"
	"net/http"
)

func ShowOpeningHandler(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, gin.H{
		"opening": opening,
	})
}
