package tests

import (
	"auth-golang-clean/internal/handler"
	"auth-golang-clean/internal/usecase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterEndpoint(t *testing.T) {

	gin.SetMode(gin.TestMode)

	mockRepo := new(MockUserRepository)

	// mock behaviour
	mockRepo.On("Create", mock.Anything).Return(nil)
	mockRepo.On("FindByEmail", "test@example.com").Return(nil, nil)

	authUsecase := usecase.NewAuthUsecase(mockRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	router := gin.Default()
	router.POST("/auth/register", authHandler.Register)

	body := map[string]string{
		"name":     "test",
		"email":    "test@example.com",
		"password": "password123",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/auth/register",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	mockRepo.AssertExpectations(t)
}
