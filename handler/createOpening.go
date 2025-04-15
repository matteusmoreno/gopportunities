package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/matteusmoreno/gopportunities/config"
	"github.com/matteusmoreno/gopportunities/schemas"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {
	logger := config.GetLogger()
	db := config.GetPostgres()

	request := CreateOpeningRequest{}

	// Bind e valida os dados da requisição
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// Se houver erro de validação, percorre os erros e retorna quais campos estão faltando
		var missingFields []string
		for _, e := range err.(validator.ValidationErrors) {
			missingFields = append(missingFields, e.Field())
		}
		// Retorna um erro detalhado com os campos ausentes
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":          "Bad Request",
			"missing_fields": missingFields,
		})
		return
	}

	// Cria o model real com base no request
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	// Salva no banco
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	// Retorna o objeto criado
	ctx.JSON(http.StatusCreated, schemas.ToOpeningResponse(opening))
}
