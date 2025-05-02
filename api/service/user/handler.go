package user

import (
	"net/http"

	"github.com/MATGILL/GIN_V2/api/types"
	"github.com/MATGILL/GIN_V2/api/utils"
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
	var userDto types.RegisterUser
	if err := utils.ParseJson(r, userDto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//verify that the user exist
	user, err := h.repository.GetUserByEmail(userDto.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	println(user.Email)
	utils.WriteJSON(w, 200, user)
}
