package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dickyanth/eco-bite-v1/service/buyer"
	"github.com/dickyanth/eco-bite-v1/service/product"
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

	buyerStore := buyer.NewStore(s.db)
	buyerHandler := buyer.NewHandler(buyerStore)
	buyerHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr,router)
}