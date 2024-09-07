package buyer

import (
	"net/http"

	"github.com/dickyanth/eco-bite-v1/types"
	"github.com/dickyanth/eco-bite-v1/utils"
	"github.com/gorilla/mux"
)

type Handler struct{
 store *types.Buyer
}

func NewHandler() *Handler{
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.handleLogin).Methods("POST")
	router.HandleFunc("/register",h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
	// get JSON Payload
	if r.Body == nil{

	}
	var payload types.RegisterBuyerPayload
	if  err := utils.ParseJSON(r,payload); err!=nil{
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if the buyer exist

}