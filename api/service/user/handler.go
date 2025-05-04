package user

import (
	"fmt"
	"net/http"

	"github.com/MATGILL/GIN_V2/api/service/auth"
	"github.com/MATGILL/GIN_V2/api/types"
	"github.com/MATGILL/GIN_V2/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	repository types.UserRepository
}

func NewHandler(repository types.UserRepository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", h.HandleLogin)
	router.POST("/register", h.HandleRegister)
}

func (h *Handler) HandleLogin(c *gin.Context) {
	var userDto types.LoginUserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(userDto); err != nil {
		err := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid dto %v", err)})
		return
	}

	user, err := h.repository.GetUserByEmail(userDto.Email)
	if err != nil || !auth.ComparePassword(user.Password, []byte(userDto.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": "token"}) // TODO: generate real token
}

func (h *Handler) HandleRegister(c *gin.Context) {
	var userDto types.RegisterUserDto

	// Gin-native JSON binding
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation avec validator
	if err := utils.Validate.Struct(userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto", "details": err.Error()})
		return
	}

	// Vérifie si l'utilisateur existe déjà
	_, err := h.repository.GetUserByEmail(userDto.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	// Hash du mot de passe
	hashedPassword, err := auth.HashPassword(userDto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Création de l'utilisateur
	err = h.repository.CreateUser(types.User{
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
