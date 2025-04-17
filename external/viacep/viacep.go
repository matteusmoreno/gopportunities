package viacep

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var viacepClient = NewClient()

func GetAddressByCEP(c *gin.Context) {
	zipcode := c.Param("cep")

	data, err := viacepClient.GetAddressByZipCode(zipcode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
