package main

import (
	"log"
	"shop-api/handler"
	"shop-api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file :: %v\n", err.Error())
	}
}

func main() {
	r := gin.New()
	r.Use(middlewares.CORSMiddleware())
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	r.POST("/new-transaction", handler.NewTransaction)
	r.GET("/search-transaction", handler.SearchTransaction)
	r.Run(":8080")
	log.Println("Server ready at port 8080")
}
