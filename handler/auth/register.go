package auth

import (
	"net/http"
	"shop-api/config"
	"shop-api/models"

	"crypto/rand"
	"encoding/base64"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	db := config.ConnectDB()
	var user models.User
	var input RegisterInput
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
		} else {
			input.Password = string(hashedPassword)
		}

		user.Id = generateRandomID()
		user.Username = input.Username
		user.Password = input.Password
		// Use struct for query value to database
		_, err = db.Exec("INSERT INTO users VALUES (?, ?, ?)", user.Id, user.Username, user.Password)
		if err != nil {
			log.Println("Error Query to database ::", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			log.Println("Success register user ::", input.Username)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Register success !!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is exist, try another one"})
		return
	}
	defer db.Close()
}

func generateRandomID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Panicln(err)
	}
	// Encode the random number to a base64 string
	encode := base64.StdEncoding.EncodeToString(b)
	replacer := strings.NewReplacer(
		"&", "",
		"-", "",
		"+", "",
		"=", "",
		"!", "",
		"/", "",
		`\`, "",
		"#", "",
		"*", "",
		"%", "",
	)
	id := replacer.Replace(encode)
	return id
}
