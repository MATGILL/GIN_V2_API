package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MATGILL/GIN_V2/api/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	userRepository := &mockUserRepository{}
	h := NewHandler(userRepository)
	r := gin.Default()
	r.POST("/register", h.HandleRegister)

	t.Run("invalid payload", func(t *testing.T) {
		body := types.RegisterUserDto{
			Firstname: "f",
			Lastname:  "l",
			Email:     "not-an-email",
			Password:  "pass",
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("successful registration", func(t *testing.T) {
		body := types.RegisterUserDto{
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "john@example.com",
			Password:  "password",
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		log.Println(resp.Code)
		assert.Equal(t, http.StatusCreated, resp.Code)
	})
}

// --- mockUserRepository ---

type mockUserRepository struct{}

func (m *mockUserRepository) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("not found")
}

func (m *mockUserRepository) GetUserById(id int) (*types.User, error) {
	return nil, fmt.Errorf("not found")
}

func (m *mockUserRepository) CreateUser(types.User) error {
	return nil
}
