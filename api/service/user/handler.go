package user

import (
	"fmt"
	"net/http"

	"github.com/MATGILL/GIN_V2/api/service/auth"
	"github.com/MATGILL/GIN_V2/api/types"
	"github.com/MATGILL/GIN_V2/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// the handler can take any dependency (for DI)
type Handler struct {
	repository types.UserRepository
}

func NewHandler(repository types.UserRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleRegister).Methods("POST")
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var userDto types.RegisterUserDto
	if err := utils.ParseJson(r, &userDto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate the dto
	if err := utils.Validate.Struct(userDto); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid dto %v", error))
		return
	}

	//verify that the user exist
	user, err := h.repository.GetUserByEmail(userDto.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", userDto.Email))
		return
	}

	//hash the password
	hashedPassword, err := auth.HashPassword(userDto.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//Create the user
	err = h.repository.CreateUser(types.User{
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}
