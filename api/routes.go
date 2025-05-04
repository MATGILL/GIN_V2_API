package api

import (
	"github.com/MATGILL/GIN_V2/api/service/user"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes enregistre toutes les routes liées aux utilisateurs.
func RegisterRoutes(router *gin.RouterGroup, handler *user.Handler) {
	router.POST("/login", handler.HandleLogin)
	router.POST("/register", handler.HandleRegister)
}
