package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MATGILL/GIN_V2/api/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandler(t *testing.T) {
	userRepository := &mockUserRepository{}
	handler := NewHandler(userRepository)

	t.Run("Should failed if the user payload is invalid", func(t *testing.T) {
		userDto := types.RegisterUserDto{
			Firstname: "fname",
			Lastname:  "Lname",
			Email:     "invalid", // not a correct email
			Password:  "password",
		}

		marshalled, _ := json.Marshal(userDto)

		request, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Unexpected status code %d, got %d", http.StatusBadRequest, recorder.Code)
		}

	})

	t.Run("Should correctly register the user", func(t *testing.T) {
		userDto := types.RegisterUserDto{
			Firstname: "fname",
			Lastname:  "Lname",
			Email:     "email@gmail.com", // not a correct email
			Password:  "password",
		}

		marshalled, _ := json.Marshal(userDto)

		request, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusCreated {
			t.Errorf("Unexpected status code %d, got %d", http.StatusCreated, recorder.Code)
		}
	})
}

// Mock the Repository
type mockUserRepository struct{}

func (m *mockUserRepository) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserRepository) GetUserById(id int) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserRepository) CreateUser(types.User) error {
	return nil
}
