package types

import "time"

type BuyerStore interface {
	GetBuyerByEmail(email string) (*Buyer, error)	
	GetBuyerById(id int) (*Buyer, error)	
	CreateBuyer(Buyer)error
}

type mockBuyerStore struct{

}

func GetBuyerByEmail(email string) (*Buyer, error) {return nil,nil}


type Buyer struct{
	Buyer_id int `json:"buyer_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone_number string `json:"phone_number"`
	Address string `json:"address"`
	Password_hash string `json:"-"`
	Created_at time.Time `json:"created_at"`
}

type RegisterBuyerPayload struct{
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"` 
	Address string `json:"address"`
	Password string `json:"password"`
}
