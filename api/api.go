package api

import (
	"database/sql"
	"log"

	"github.com/MATGILL/GIN_V2/api/service/user"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := gin.Default()

	// Appliquer le middleware ici si nécessaire
	s.applyMiddleware(router)

	// Configuration des routes de l'API
	api := router.Group("/api/v1")
	s.configureRoutes(api)

	log.Println("Listening on", s.addr)
	return router.Run(s.addr)
}

func (s *ApiServer) configureRoutes(router *gin.RouterGroup) {
	// Initialisation des handlers
	userRepository := user.NewRepository(s.db)
	userHandler := user.NewHandler(userRepository)

	// Inscription des routes avec la fonction RegisterRoutes
	RegisterRoutes(router, userHandler)
}

func (s *ApiServer) applyMiddleware(router *gin.Engine) {
	// Exemples de middleware : Log, Auth, etc.
	// router.Use(middleware.SomeMiddleware())

	// Par exemple, un middleware de log
	router.Use(func(c *gin.Context) {
		log.Printf("Requeeeeest: %s %s", c.Request.Method, c.Request.URL)
		c.Next()
	})
}
