package auth

import (
	"net/http"
	"shop-api/config"
	"shop-api/models"
	"shop-api/utils"

	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullname" binding:"required"`
}

func Register(c *gin.Context) {
	db := config.ConnectDB()
	var user models.User
	var input registerInput
	var profile models.Profile

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var checkUsername error
	checkUsername = db.QueryRow("SELECT username FROM users WHERE username = ?", input.Username).Scan(&input.Username)
	if checkUsername != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error hashing password ::", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
			return
		}
		input.Password = string(hashedPassword)

		user.Id = utils.GenerateRandomID(8)
		user.Username = input.Username
		user.Password = input.Password
		// Use struct for query value to database
		_, err = db.Exec("INSERT INTO users VALUES (?, ?, ?)", user.Id, user.Username, user.Password)
		if err != nil {
			log.Println("Error Query to database ::", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error input data"})
			return
		}

		profile.Id = utils.GenerateRandomID(8)
		profile.UserId = user.Id
		profile.FullName = input.FullName
		_, err = db.Exec("INSERT INTO profiles (id, user_id, full_name) VALUES (?, ?, ?)", profile.Id, profile.UserId, profile.FullName)
		if err != nil {
			log.Println("Error Query to database ::", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error input data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Register success !!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is exist, try another one"})
		return
	}
	defer db.Close()
}
