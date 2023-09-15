package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type CartUseCase interface {
	// AddToCart adds a product to the user's shopping cart.
	AddToCart(userID, ProductID int) error

	// ViewCart retrieves the items in the user's shopping cart.
	ViewCart(userID int) (response.CartItems, error)

	// RemoveFromCart removes a product from the user's shopping cart.
	RemoveFromCart(userID, productID int) error

	// IncrementQuantity increases the quantity of a product in the user's cart.
	IncrementQuantity(userID, productID int) error

	// DecrementQuantity decreases the quantity of a product in th user's cart.
	DecrementQuantity(userID, productID int) error

	// DeleteUserCart deletes the entire shopping cart of a user.
	DeleteUserCart(userID int) error
}
