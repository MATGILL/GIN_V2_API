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
	api := router.Group("/api/v1")

	userRepository := user.NewRepository(s.db)
	userHandler := user.NewHandler(userRepository)
	userHandler.RegisterRoutes(api)

	log.Println("Listening on", s.addr)
	return router.Run(s.addr)
}
