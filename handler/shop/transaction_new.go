package shop

import (
	"fmt"
	"net/http"
	"shop-api/config"
	"shop-api/models"
	"shop-api/utils"

	"github.com/gin-gonic/gin"
)

type newTransactionInput struct {
	Menu    string `json:"menu"`
	Price   int    `json:"price"`
	Qty     int    `json:"qty"`
	Payment string `json:"payment"`
	Total   int    `json:"total"`
}

func NewTransaction(c *gin.Context) {
	db := config.ConnectDB()
	transaction_input := newTransactionInput{}
	transaction := models.Transaction{}

	if err := c.BindJSON(&transaction_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, errs := config.ExtractTokenID(c)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}
	user_id := fmt.Sprint(uid)

	transaction.Id = utils.GenerateRandomID(11)
	transaction.Customer_id = user_id
	transaction.Menu = transaction_input.Menu
	transaction.Price = transaction_input.Price
	transaction.Qty = transaction_input.Qty
	transaction.Payment = transaction_input.Payment
	transaction.Total = transaction_input.Total
	_, err := db.Exec(
		"INSERT INTO transactions VALUES (?, ?, ?, ?, ?, ?, ?, NOW())",
		transaction.Id, transaction.Customer_id, transaction.Menu,
		transaction.Price, transaction.Qty, transaction.Payment, transaction.Total,
	)
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction success"})
}
