package buyer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dickyanth/eco-bite-v1/types"
	"github.com/gorilla/mux"
)

func TestBuyerService(t *testing.T){
	buyerStore := &mockBuyerStore{}
handler := NewHandler(buyerStore)

	t.Run("should fail if the buyer payload is invalid", func (t *testing.T){
		payload := types.RegisterBuyerPayload{
			Name:"user",
			Email:"invalid",
			PhoneNumber:"123",
			Address:"123",
			Password:"123",
		}
		marshalled, _ := json.Marshal(payload)

		req, err:=http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		
		if rr.Code != http.StatusBadRequest{
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("Should correctly registered user", func (t *testing.T) {
		payload := types.RegisterBuyerPayload{
			Name:"user",
			Email:"asd@gmail.com",
			PhoneNumber:"123",
			Address:"123",
			Password:"123",
		}
		marshalled, _ := json.Marshal(payload)

		req, err:=http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		
		if rr.Code != http.StatusCreated{
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockBuyerStore struct {}

func (m *mockBuyerStore) GetBuyerByEmail(email string) (*types.Buyer,error){
	return nil, fmt.Errorf("User not found")
}

func (m *mockBuyerStore) GetBuyerByID(id int) (*types.Buyer,error){
	return nil,nil
}

func (m *mockBuyerStore) CreateBuyer(types.Buyer) error{
	return nil
}