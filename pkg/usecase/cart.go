package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type CartUseCase struct {
	cartRepo   interfaces.CartRepository
	coupenRepo interfaces.CouponRepository
}

func NewCartUseCase(cartUseCase interfaces.CartRepository, coupenUseCase interfaces.CouponRepository) services.CartUseCase {
	return &CartUseCase{
		cartRepo:   cartUseCase,
		coupenRepo: coupenUseCase,
	}
}

func (cu *CartUseCase) AddToCart(userID int, productID int) error {
	CartItem, err := cu.cartRepo.GetCartItem(userID, productID)
	fmt.Println("user id :", userID)
	if err != nil {
		return fmt.Errorf("Failed to add to cart : %s", err)
	}
	if CartItem.ID != 0 {
		err = cu.IncrementQuantity(userID, productID)
		if err != nil {
			return fmt.Errorf("Failed %s", err)
		}
		return nil
	}

	CartItem, err = cu.cartRepo.AddToCart(userID, productID)
	if err != nil || CartItem.ID == 0 {
		return fmt.Errorf("Add to cart failed:%s", err)
	}
	return nil
}
func (cu *CartUseCase) ViewCart(userID int) (response.CartItems, error) {
	CartInfo, err := cu.cartRepo.ViewCart(userID)

	var CartItems response.CartItems
	for _, item := range CartInfo {
		CartItems.Cart = append(CartItems.Cart, item)
		CartItems.Total += float32(item.Qty) * float32(item.Price)
	}

	CouponDetails, err := cu.coupenRepo.CheckForAppliedCoupon(userID)
	if err != nil {
		return response.CartItems{}, fmt.Errorf("Failed to fetch coupon details")
	}
	var DiscountPrize float64
	if CouponDetails.ID != 0 && CouponDetails.CouponID != 0 {
		Coupon, err := cu.coupenRepo.FindCouponById(CouponDetails.CouponID)
		fmt.Println(Coupon)
		if err != nil {
			return response.CartItems{}, fmt.Errorf("Failed to find coupon :%s", err)
		}
		fmt.Println(Coupon.ValidTill)
		fmt.Println(helper.IsCouponValid(Coupon.ValidTill))
		if !helper.IsCouponValid(Coupon.ValidTill) {
			return response.CartItems{}, fmt.Errorf("Expired coupon")
		}
		if CartItems.Total > float32(Coupon.MinOrderValue) {
			DiscountPrize = (float64(CartItems.Total) * Coupon.DiscountPercent) / 100
			if DiscountPrize > Coupon.DiscountMaxAmount {
				DiscountPrize = Coupon.DiscountMaxAmount
			}
		}
	}
	CartItems.Discount = float32(DiscountPrize)
	CartItems.Total = CartItems.Total - float32(DiscountPrize)

	fmt.Println("CART ITEMS PRIZES :::", CartItems.Discount, CartItems.Total)
	return CartItems, err

}

func (cu *CartUseCase) RemoveFromCart(userID int, productID int) error {
	CartItem, err := cu.cartRepo.RemoveFromCart(userID, productID)
	fmt.Println("CART ITEM:::", CartItem)
	fmt.Println(userID, productID)
	if err != nil {
		return fmt.Errorf("Remove from cart failed:%s", err)

	}
	return nil
}
func (cu *CartUseCase) IncrementQuantity(userID int, productID int) error {
	CartItem, err := cu.cartRepo.GetCartItem(userID, productID)
	if err != nil || CartItem.ID == 0 {
		return fmt.Errorf("Quantity updation  failed:%s", err)

	}
	qty := CartItem.Qty
	newQty := qty + 1

	CartItem, err = cu.cartRepo.IncrementQuantity(newQty, userID, productID)
	if err != nil || CartItem.ID == 0 || newQty != CartItem.Qty {
		return fmt.Errorf("Quantity updation failed : %s", err)
	}
	return nil
}

func (cu *CartUseCase) DecrementQuantity(userID int, productID int) error {
	CartItem, err := cu.cartRepo.GetCartItem(userID, productID)
	if err != nil || CartItem.ID == 0 {
		return fmt.Errorf("Quantity updation  failed:%s", err)

	} else if CartItem.Qty == 1 {
		return nil
	}

	qty := CartItem.Qty
	newQty := qty - 1

	CartItem, err = cu.cartRepo.DecrementQuantity(newQty, userID, productID)
	if err != nil || CartItem.ID == 0 || newQty != CartItem.Qty {
		return fmt.Errorf("Quantity updation failes : %s", err)
	}
	return nil
}

func (cu *CartUseCase) DeleteUserCart(userID int) error {
	_, err := cu.cartRepo.DeleteCart(userID)
	if err != nil {
		return fmt.Errorf("failed to delete user cart  : %s", err)
	}
	// if DeletedCart.ID == 0 {
	// 	return fmt.Errorf("failed to delete user cart  :showing deleted cart id = %d", DeletedCart.ID)
	// }
	return nil

}
