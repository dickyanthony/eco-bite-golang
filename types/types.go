package types

import "time"

type BuyerStore interface {
	GetBuyerByEmail(email string) (*Buyer, error)	
	GetBuyerByID(id int) (*Buyer, error)	
	CreateBuyer(Buyer)error
}

type ProductStore interface{
	GetProducts() ([]Product, error)
	GetProductByIds(ps []int)([]Product, error)
	UpdateProduct(Product) error
}

type OrderStore interface{
	CreateOrder(Order)(int, error)
	CreateOrderItem(OrderItem) error
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type Order struct{
	Id int `json:"id"`
	BuyerId int `json:"buyerId"`
	Total float64 `json:"total"`
	Status string `json:"status"`
	Address string `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct{
	Id int `json:"id"`
	OrderId int `json:"orderId"`
	ProductId int `json:"productId"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct{
	Id int `json:"id"`
	Seller_id int `json:"sellerId"`
	Image string `json:"image"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	ExpiryDate time.Time `json:"expiryDate"`
	CreatedAt time.Time `json:"createdAt"`
}

type mockBuyerStore struct{

}

func GetBuyerByEmail(email string) (*Buyer, error) {return nil,nil}

type Buyer struct{
	BuyerId int `json:"buyer_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
	PasswordHash string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
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

type CartItem struct{
	ProductId int
	Quantity int
}

type CartCheckoutPayload struct{
	Items []CartItem `json:"items", validate:"required"`
}