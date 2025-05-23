package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MATGILL/GIN_V2/api/service/user"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	//Look in depth here for understand di and add a service layer
	userRepository := user.NewRepository(s.db)
	userHandler := user.NewHandler(userRepository)

	userHandler.RegisterRoutes(subRouter)

	log.Println("Listenning on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
