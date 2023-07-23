package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type CartUseCase interface {
	AddToCart(userID int, ProductID int) error
	ViewCart(userID int) (response.CartItems, error)
	RemoveFromCart(userID int, productID int) error
	IncrementQuantity(userID int, productID int) error
	DecrementQuantity(userID int, productID int) error
	DeleteUserCart(userID int) error
}
