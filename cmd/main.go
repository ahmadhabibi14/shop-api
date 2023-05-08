package main

import (
	"log"
	"shop-api/handler/auth"
	"shop-api/handler/setting/user"
	"shop-api/handler/shop"
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
	// Public API
	public := r.Group("/api")
	public.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSMiddleware())
	public.POST("/login", auth.Login)
	public.POST("/register", auth.Register)

	// Protected API, auth is required
	shop_api := r.Group("/api/shop")
	shop_api.Use(gin.Recovery(), middlewares.Logger(), middlewares.AuthJWT())
	shop_api.POST("/new-transaction", shop.NewTransaction)
	shop_api.GET("/search-transaction", shop.SearchTransaction)

	user_setting := r.Group("/api/setting")
	user_setting.Use(gin.Recovery(), middlewares.Logger(), middlewares.AuthJWT())
	user_setting.POST("/avatar-image", user.EditAvatar)

	// SERVER READY !!
	r.Run(":8080")
	log.Println("Server ready at port 8080")
}
