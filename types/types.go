package types

import "time"

type BuyerStore interface {
	GetBuyerByEmail(email string) (*Buyer, error)	
	GetBuyerByID(id int) (*Buyer, error)	
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
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"` 
	Address string `json:"address" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}
