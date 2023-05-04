package handler

import (
	"net/http"
	"shop-api/config"
	"shop-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user models.User
	// Convert request body to struct
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := LoginCheck(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func LoginCheck(username, password string) (string, error) {
	db := config.ConnectDB()
	user := models.User{}
	var err error

	// Check if username is exist
	err = db.QueryRow("SELECT Username, Password FROM User WHERE Username = ?", username).Scan(&user.Username, &user.Password)
	if err != nil {
		return "", err
	}
	defer db.Close()
	// Check if password is exist
	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", nil
	}

	message := "Sucess login !"
	return message, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
