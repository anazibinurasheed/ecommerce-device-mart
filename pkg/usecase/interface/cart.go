package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type CartUseCase interface {
	AddToCart(userID, ProductID int) error
	ViewCart(userID int) (response.CartItems, error)
	RemoveFromCart(userID, productID int) error
	IncrementQuantity(userID, productID int) error
	DecrementQuantity(userID, productID int) error
	DeleteUserCart(userID int) error
}
