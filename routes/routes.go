package routes

import (
	"auth-golang-clean/internal/handler"
	"auth-golang-clean/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler *handler.AuthHandler) {

	auth := r.Group("/auth")

	{
		auth.POST("/register", authHandler.Register)

		auth.POST("/login", authHandler.Login)
	}

	protected := r.Group("/api")

	protected.Use(middleware.JWTMiddleware())

	{
		protected.GET("/profile", func(c *gin.Context) {

			c.JSON(200, gin.H{

				"message": "protected route",
			})
		})
	}
}
