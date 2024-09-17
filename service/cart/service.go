package cart

import (
	"fmt"

	"github.com/dickyanth/eco-bite-v1/types"
)

func getCartItemIds(items []types.CartItem)([]int, error){
	productIds := make([]int, len(items))
	for i, item := range items{
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductId)
		}
		productIds[i] = item.ProductId
	}

	return productIds,nil

}

func (h *Handler) createOrder(ps []types.Product, items []types.CartItem, buyerId int) (int, float64, error){
	productMap := make(map[int]types.Product)
	for _, product := range ps{
		productMap[product.Id] = product
	}

	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}

	totalPrice := calculateTotalPrice(items, productMap)

	for _, item := range items{
		product := productMap[item.ProductId]
		product.Quantity -= item.Quantity 
		h.productStore.UpdateProduct(product)
	}

	orderId, err := h.store.CreateOrder(types.Order{
		BuyerId: buyerId,
		Total:totalPrice,
		Status: "pending",
		Address: "Some Address",
	})

for _, item := range items{
	h.store.CreateOrderItem(types.OrderItem{
		OrderId: orderId,
		ProductId: item.ProductId,
		Quantity: item.Quantity,
		Price: productMap[item.ProductId].Price,
	})
}

	return orderId, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error{
	if len(cartItems) == 0{
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems{
		product, ok := products[item.ProductId]
		if !ok{
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductId)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItem []types.CartItem, products map[int]types.Product) float64{
	var total float64

	for _, item := range cartItem{
		product := products[item.ProductId]
		total += product.Price * float64(item.Quantity)
	}

	return total
}

