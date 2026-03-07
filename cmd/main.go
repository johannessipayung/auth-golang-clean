package main

import (
	"auth-golang-clean/config"
	"auth-golang-clean/internal/entity"
	"auth-golang-clean/internal/handler"
	"auth-golang-clean/internal/repository"
	"auth-golang-clean/internal/usecase"
	"auth-golang-clean/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitLogger()

	db, err := config.ConnectDatabase()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo)

	authHandler := handler.NewAuthHandler(authUsecase)

	r := gin.Default()

	routes.SetupRoutes(r, authHandler)

	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
