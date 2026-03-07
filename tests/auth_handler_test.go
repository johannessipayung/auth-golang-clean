package tests

import (
	"auth-golang-clean/internal/handler"
	"auth-golang-clean/internal/usecase"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterEndpoint(t *testing.T) {

	gin.SetMode(gin.TestMode)

	mockRepo := new(MockUserRepository)

	authUsecase := usecase.NewAuthUsecase(mockRepo)

	authHandler := handler.NewAuthHandler(authUsecase)

	router := gin.Default()

	router.POST("/auth/register", authHandler.Register)

	body := []byte(`{
		"username":"test",
		"email":"test@mail.com",
		"password":"123456"
	}`)

	req, _ := http.NewRequest(
		"POST",
		"/auth/register",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", w.Code)
	}
}
