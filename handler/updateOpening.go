package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/matteusmoreno/gopportunities/config"
	"github.com/matteusmoreno/gopportunities/schemas"
	"net/http"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	db := config.GetPostgres()
	logger := config.GetLogger()

	var request UpdateOpeningRequest

	// Bind do JSON da requisição
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Busca a vaga existente no banco
	var opening schemas.Opening
	if err := db.First(&opening, request.ID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Opening not found"})
		return
	}

	// Atualiza apenas os campos que foram enviados
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary != 0 {
		opening.Salary = request.Salary
	}

	// Salva as mudanças no banco
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update opening"})
		return
	}

	// Resposta de sucesso
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Opening updated successfully",
		"data":    opening,
	})
}
