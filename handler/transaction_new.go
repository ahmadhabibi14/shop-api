package handler

import (
	"log"
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
	_, err := db.Exec(
		"INSERT INTO Transaction (id, customer_id, menu, price, qty, payment, total) VALUES (?, ?, ?, ?, ?, ?, ?",
		transaction.Id, transaction.Customer_id, transaction.Menu,
		transaction.Price, transaction.Qty, transaction.Payment, transaction.Total,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err.Error())
	}
	defer db.Close()
	c.JSON(http.StatusOK, gin.H{"message": "Transaction success"})
}
