package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/matteusmoreno/gopportunities/schemas"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening
	if err := db.Find(&openings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	//to opening response list
	var openingResponses []schemas.OpeningResponse
	for _, opening := range openings {
		openingResponses = append(openingResponses, schemas.ToOpeningResponse(opening))
	}

	ctx.JSON(http.StatusOK, openingResponses)
}
