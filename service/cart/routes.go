package cart

import (
	"fmt"
	"net/http"

	"github.com/dickyanth/eco-bite-v1/service/auth"
	"github.com/dickyanth/eco-bite-v1/types"
	"github.com/dickyanth/eco-bite-v1/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)


type Handler struct {
	store types.OrderStore
	productStore types.ProductStore
	buyerStore types.BuyerStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, buyerStore types.BuyerStore) *Handler {
 return &Handler{store,productStore,buyerStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/cart/checkout",auth.WithJWTAuth(h.handleCheckout, h.buyerStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request){
	buyerId := auth.GetUserIdFromContext(r.Context())
	var cart  types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}


	productIds, err:= getCartItemIds(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return 
	}
	ps, err := h.productStore.GetProductByIds(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return 
	}

	orderId, totalPrice, err := h.createOrder(ps, cart.Items, buyerId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return 
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"total_price": totalPrice,
		"order_id": orderId,
	})
}