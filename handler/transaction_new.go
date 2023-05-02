package handler

import (
	"net/http"
	"shop-api/config"
	"shop-api/models"

	"github.com/gin-gonic/gin"
)

func NewTransaction(c *gin.Context) {
	transaction := models.Transaction{}
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := config.ConnectDB()
	_, err := db.Exec("INSERT INTO Transaction")
}
