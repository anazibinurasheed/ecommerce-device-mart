package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type CartRepository interface {
	AddToCart(userID int, ProductID int) (response.Cart, error)
	ViewCart(userID int) ([]response.Cart, error)
	RemoveFromCart(userID int, productID int) (response.Cart, error)
	IncrementQuantity(qty int, userID int, productID int) (response.Cart, error)
	DecrementQuantity(qty int, userID int, productID int) (response.Cart, error)
	GetCartItem(userID int, productID int) (response.Cart, error)
	DeleteCart(userID int) (response.Cart, error)
}
