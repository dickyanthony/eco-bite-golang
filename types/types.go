package types

import "time"

type BuyerStore interface {
	GetBuyerByEmail(email string) (*Buyer, error)	
	GetBuyerByID(id int) (*Buyer, error)	
	CreateBuyer(Buyer)error
}

type ProductStore interface{
	GetProducts() ([]Product, error)
}

type Product struct{
	Id int `json:"id"`
	Seller_id int `json:"seller_id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	Expiry_date time.Time `json:"expiry_date"`
	Created_at time.Time `json:"created_at"`
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

type LoginBuyerPayload struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}
