package auth

import (
	"net/http"
	"shop-api/config"
	"shop-api/models"

	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user models.User
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Username = input.Username
	user.Password = input.Password
	token, err := LoginCheck(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": token})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func LoginCheck(username, password string) (string, error) {
	db := config.ConnectDB()
	user := models.User{}
	var err error
	err = db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return "username not found", err
	}
	defer db.Close()
	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "incorrect password", err
	}
	fmt.Println(user.Id, user.Username, user.Password)
	token, err := config.GenerateJWT(user.Id)
	if err != nil {
		return "error generate JWT", err
	}
	return token, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
