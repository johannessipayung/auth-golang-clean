package handler

import (
	"auth-golang-clean/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(uc usecase.AuthUsecase) *AuthHandler {

	return &AuthHandler{uc}
}

type AuthRequest struct {
	Username string `json:"username"`

	Email string `json:"email"`

	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := h.authUsecase.Register(req.Username, req.Email, req.Password)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, gin.H{"message": "user registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, refresh, err := h.authUsecase.Login(req.Email, req.Password)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, gin.H{

		"token":         token,
		"refresh_token": refresh,
	})
}
