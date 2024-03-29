package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahmadhabibi14/shop-api/app"
	"github.com/ahmadhabibi14/shop-api/controller"
	"github.com/ahmadhabibi14/shop-api/helper"
	"github.com/ahmadhabibi14/shop-api/middleware"
	"github.com/ahmadhabibi14/shop-api/repository"
	"github.com/ahmadhabibi14/shop-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
}

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    os.Getenv("WEB_DOMAIN"),
		Handler: authMiddleware,
	}
}

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)

	server := NewServer(authMiddleware)

	log.Println("Server ready at port", os.Getenv("WEB_PORT"))
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
