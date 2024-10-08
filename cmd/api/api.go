package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamrubayet/ecom/service/user"
	"github.com/rs/zerolog/log"
)

type APISERVER struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APISERVER {
	return &APISERVER{
		addr: addr,
		db:   db,
	}

}

func (s *APISERVER) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	log.Print("Server is running on port ", s.addr)

	return http.ListenAndServe(s.addr, router)

}
