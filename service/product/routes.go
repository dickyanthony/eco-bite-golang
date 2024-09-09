package product

import (
	"net/http"

	"github.com/dickyanth/eco-bite-v1/types"
	"github.com/dickyanth/eco-bite-v1/utils"
	"github.com/gorilla/mux"
)

type Handler struct{
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler{
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request){
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)

}