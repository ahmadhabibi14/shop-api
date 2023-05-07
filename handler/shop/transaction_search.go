package shop

import (
	"net/http"

	"shop-api/config"
	"shop-api/models"

	"github.com/gin-gonic/gin"
)

func SearchTransaction(c *gin.Context) {
	db := config.ConnectDB()
	var transactions []models.Transaction

	user_id, errs := config.ExtractTokenID(c)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}
	var rows, err = db.Query(
		"SELECT transactions.* FROM transactions JOIN users on transactions.customer_id = users.id WHERE transactions.customer_id = ?",
		user_id,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no transaction yet"})
		return
	}
	for rows.Next() {
		var each = models.Transaction{}
		var err = rows.Scan(
			&each.Id, &each.Customer_id, &each.Menu, &each.Price,
			&each.Qty, &each.Payment, &each.Total, &each.Created_at,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot migrate to struct"})
			return
		}
		transactions = append(transactions, each)
	}
	defer db.Close()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, transactions)
}
