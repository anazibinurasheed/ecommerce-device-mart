package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type CartUseCase struct {
	cartRepo   interfaces.CartRepository
	couponRepo interfaces.CouponRepository
}

func NewCartUseCase(cartUseCase interfaces.CartRepository, couponUseCase interfaces.CouponRepository) services.CartUseCase {
	return &CartUseCase{
		cartRepo:   cartUseCase,
		couponRepo: couponUseCase,
	}
}

func (cu *CartUseCase) AddToCart(userID, productID int) error {
	cartItem, err := cu.cartRepo.GetCartItem(userID, productID)
	if err != nil {
		return fmt.Errorf("Failed to add to cart :%s", err)
	}

	if cartItem.ID != 0 {
		err = cu.IncrementQuantity(userID, productID)
		if err != nil {
			return fmt.Errorf("Failed to increment quantity :%s", err)
		}
		return nil
	}

	cartItem, err = cu.cartRepo.AddToCart(userID, productID)
	if err != nil || cartItem.ID == 0 {
		return fmt.Errorf("Add to cart failed:%s", err)
	}
	return nil
}

func (cu *CartUseCase) ViewCart(userID int) (response.CartItems, error) {
	cart, err := cu.cartRepo.ViewCart(userID)
	if err != nil {
		return response.CartItems{}, fmt.Errorf("Failed to fetch user cart :%s", err)
	}

	var cartItems response.CartItems
	for _, item := range cart {
		cartItems.Cart = append(cartItems.Cart, item)
		cartItems.Total += float32(item.Qty) * float32(item.Price)
	}

	couponDetails, err := cu.couponRepo.CheckAppliedCoupon(userID)
	if err != nil {
		return response.CartItems{}, fmt.Errorf("Failed to fetch coupon details")
	}

	var discountPrize float64

	if couponDetails.ID != 0 && couponDetails.CouponID != 0 {
		Coupon, err := cu.couponRepo.FindCouponByID(couponDetails.CouponID)
		fmt.Println(Coupon)
		if err != nil {
			return response.CartItems{}, fmt.Errorf("Failed to find coupon :%s", err)
		}

		if !helper.IsCouponValid(Coupon.ValidTill) {
			return response.CartItems{}, fmt.Errorf("Expired coupon")
		}
		if cartItems.Total > float32(Coupon.MinOrderValue) {
			discountPrize = (float64(cartItems.Total) * Coupon.DiscountPercent) / 100
			if discountPrize > Coupon.DiscountMaxAmount {
				discountPrize = Coupon.DiscountMaxAmount
			}
		}
	}

	cartItems.Discount = float32(discountPrize)
	cartItems.Total = cartItems.Total - float32(discountPrize)

	return cartItems, err
}

func (cu *CartUseCase) RemoveFromCart(userID, productID int) error {
	cartItem, err := cu.cartRepo.RemoveFromCart(userID, productID)
	if err != nil {
		return fmt.Errorf("Remove from cart failed :%s", err)
	}
	if cartItem.ID == 0 {
		return fmt.Errorf("Failed to verify removed product")
	}
	return nil
}

func (cu *CartUseCase) IncrementQuantity(userID, productID int) error {
	cartItem, err := cu.cartRepo.GetCartItem(userID, productID)
	if err != nil || cartItem.ID == 0 {
		return fmt.Errorf("Quantity updation failed :%s", err)
	}

	qty := cartItem.Qty
	newQty := qty + 1

	cartItem, err = cu.cartRepo.IncrementQuantity(newQty, userID, productID)
	if err != nil || cartItem.ID == 0 || newQty != cartItem.Qty {
		return fmt.Errorf("Quantity updation failed : %s", err)
	}
	return nil
}

func (cu *CartUseCase) DecrementQuantity(userID, productID int) error {
	cartItem, err := cu.cartRepo.GetCartItem(userID, productID)
	if err != nil || cartItem.ID == 0 {
		return fmt.Errorf("Quantity updation failed :%s", err)

	} else if cartItem.Qty == 1 {
		return nil
	}

	qty := cartItem.Qty
	newQty := qty - 1

	cartItem, err = cu.cartRepo.DecrementQuantity(newQty, userID, productID)
	if err != nil || cartItem.ID == 0 || newQty != cartItem.Qty {
		return fmt.Errorf("Quantity updation failed :%s", err)
	}
	return nil
}

func (cu *CartUseCase) DeleteUserCart(userID int) error {
	_, err := cu.cartRepo.DeleteCart(userID)
	if err != nil {
		return fmt.Errorf("failed to delete user cart  : %s", err)
	}

	return nil

}
