package main

import (
	"auth-golang-clean/config"
	"auth-golang-clean/internal/entity"
	"auth-golang-clean/internal/handler"
	"auth-golang-clean/internal/repository"
	"auth-golang-clean/internal/usecase"
	"auth-golang-clean/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitLogger()

	db, err := config.ConnectDatabase()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	db.AutoMigrate(&entity.User{})

	userRepo := repository.NewUserRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo)

	authHandler := handler.NewAuthHandler(authUsecase)

	r := gin.Default()

	routes.SetupRoutes(r, authHandler)

	r.Run(":9090")
}
