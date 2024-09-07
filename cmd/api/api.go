package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dickyanth/eco-bite-v1/service/buyer"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string,db *sql.DB) *APIServer {
	return &APIServer{
		addr,db,
	}
}

func (s *APIServer) Run() error{
	router:=mux.NewRouter()
	subrouter:=router.PathPrefix("/api/v1/").Subrouter()

	buyerHandler := buyer.NewHandler()
	buyerHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr,router)
}