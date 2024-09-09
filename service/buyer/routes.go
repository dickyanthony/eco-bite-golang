package buyer

import (
	"fmt"
	"net/http"

	"github.com/dickyanth/eco-bite-v1/config"
	"github.com/dickyanth/eco-bite-v1/service/auth"
	"github.com/dickyanth/eco-bite-v1/types"
	"github.com/dickyanth/eco-bite-v1/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct{
 	store types.BuyerStore
}

func NewHandler(store types.BuyerStore) *Handler{
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.handleLogin).Methods("POST")
	router.HandleFunc("/register",h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){
	// get JSON Payload
	if r.Body == nil{

	}
	var payload types.LoginBuyerPayload
	if  err := utils.ParseJSON(r, &payload); err!=nil{
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate payload
	if err:=utils.Validate.Struct(payload); err != nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest,fmt.Errorf("invalid payload %v", errors))
		return 
	}

	u, err := h.store.GetBuyerByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePassword(u.Password_hash, []byte(payload.Password)){
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token,err := auth.CreateJWT(secret, u.Buyer_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return 
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token":token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
	// get JSON Payload
	if r.Body == nil{

	}
	var payload types.RegisterBuyerPayload
	if  err := utils.ParseJSON(r, &payload); err!=nil{
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate payload
	if err:=utils.Validate.Struct(payload); err != nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest,fmt.Errorf("invalid payload %v", errors))
		return 
	}

	// check if the buyer exist
	_, err := h.store.GetBuyerByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword,err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return 
	}

	// If not exist 
	err = h.store.CreateBuyer(types.Buyer{
		Name:payload.Name,
		Email: payload.Email,
		Phone_number: payload.PhoneNumber,
		Address: payload.Address,
		Password_hash: hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}