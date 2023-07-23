package usecase

import (
	"fmt"
	"sync"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type orderUseCase struct {
	userRepo    interfaces.UserRepository
	cartUseCase services.CartUseCase
	paymentRepo interfaces.PaymentRepository
	orderRepo   interfaces.OrderRepository
	couponRepo  interfaces.CouponRepository
	walletRepo  interfaces.WalletRepository
}

//some times error will invoke from here
func NewOrderUseCase(UserUseCase interfaces.UserRepository, CartUseCase services.CartUseCase, paymentUseCase interfaces.PaymentRepository, OrderUseCase interfaces.OrderRepository, CouponUseCase interfaces.CouponRepository) services.OrderUseCase {
	return &orderUseCase{
		userRepo:    UserUseCase,
		cartUseCase: CartUseCase,
		paymentRepo: paymentUseCase,
		orderRepo:   OrderUseCase,
		couponRepo:  CouponUseCase,
	}
}

//safe for concurrent use by multiple goroutines without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.

func (ou *orderUseCase) CheckOutDetails(userID int) (response.CheckOut, error) {
	var CheckOutDetails response.CheckOut

	addresses, err := ou.userRepo.GetAllUserAddresses(userID)
	if err != nil {
		return response.CheckOut{}, fmt.Errorf("Failed to retrieve checkout details %s", err)
	}
	if len(addresses) == 0 {
		return response.CheckOut{}, fmt.Errorf("User have no address")
	}

	cartItems, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return response.CheckOut{}, fmt.Errorf("Failed to retrieve cartitems %s", err)
	}

	paymentMethods, err := ou.paymentRepo.GetPaymentMethods()
	if err != nil || paymentMethods == nil {
		return response.CheckOut{}, fmt.Errorf("Failed to get payment methods %s", err)
	}
	CheckOutDetails.Address = addresses
	CheckOutDetails.Cart = cartItems.Cart
	CheckOutDetails.Total = cartItems.Total
	CheckOutDetails.Discount = cartItems.Discount
	CheckOutDetails.PaymentOptions = paymentMethods
	fmt.Println("CHECKOUT DETAILS:::::", CheckOutDetails)

	return CheckOutDetails, nil
}

func (ou *orderUseCase) GetRazorPayDetails(userID int) (response.PaymentDetails, error) {
	CartData, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to retrieve cartitems %s", err)
	}

	UserData, err := ou.userRepo.FindUserById(userID)
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to find user  %s", err)
	}
	RazorPayOrderId, err := helper.MakeRazorPayPaymentId(int(CartData.Total * 100))
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to get razorpay id %s", err)
	}
	return response.PaymentDetails{
		Username:        UserData.UserName,
		RazorPayOrderId: RazorPayOrderId,
		Amount:          int(CartData.Total),
	}, nil

}
func (ou *orderUseCase) VerifyRazorPayPayment(signature string, razorpayOrderId string, paymentId string) error {
	err := helper.VerifyRazorPayPayment(signature, razorpayOrderId, paymentId)
	if err != nil {
		return fmt.Errorf("Failed payment not success : %s", err)
	}
	return nil
}

var sessionPayment sync.Map

func (ou *orderUseCase) PaymentOption(userID int, methodID int) {

	sessionPayment.Store(userID, methodID)
	// sessionPayment.Range(func(key, value interface{}) bool {
	// 	fmt.Printf("Key: %v, Value: %v\n", key, value)
	// 	return true
	// })
	return
}

func (ou *orderUseCase) ConfirmedOrder(userID int, paymentMethodID int) error {

	address, err := ou.userRepo.FindDefaultAddressById(userID)
	if err != nil || address.ID == 0 {
		return fmt.Errorf("Failed to find default address : %s", err)
	}
	addressID := address.ID
	// methodId, ok := sessionPayment.Load(userID)
	// if !ok {
	// 	return fmt.Errorf("Failed to retrieve paymentMethodId from paymentSession : %s", err)
	// }
	// paymentMethodID := methodId.(int)

	cartData, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return fmt.Errorf("Failed to get Cart data :  %s", err)
	}
	CouponDetails, err := ou.couponRepo.FindAppliedCouponByUserId(userID)
	if err != nil {
		return fmt.Errorf("Failed to retrieve the coupon tracking details  ;%s", err)
	}

	if CouponDetails.CouponID != 0 {
		UpdatedCouponTracking, err := ou.couponRepo.UpdateCouponUsage(userID)
		if err != nil {
			return fmt.Errorf("Failed to update coupon usage :%s", err)
		}
		if UpdatedCouponTracking.ID == 0 {
			return fmt.Errorf("Failed to verify inserted coupon record")
		}

	}
	Coupon, err := ou.couponRepo.FindCouponById(CouponDetails.CouponID)
	if !helper.IsCouponValid(Coupon.ValidTill) {
		return fmt.Errorf("Failed applied Coupon is expired")
	}
	status, err := ou.orderRepo.GetStatusPending()
	if err != nil {
		return fmt.Errorf("Failed to get order status :%s", err)
	}
	statusID := status.ID

	for _, productData := range cartData.Cart {
		newOrderLine, err := ou.orderRepo.InsertOrderLine(userID, int(productData.ProductID), int(addressID), productData.Qty, productData.Price, paymentMethodID, int(statusID), CouponDetails.CouponID)
		if err != nil || newOrderLine.ID == 0 {
			return fmt.Errorf("Failed to insert order line : %s", err)
		}

	}

	err = ou.cartUseCase.DeleteUserCart(userID)
	if err != nil {
		return fmt.Errorf("Failed to delete user cart :%s", err)
	}

	return nil
}

func (ou *orderUseCase) GetUserOrderHistory(userID, page, count int) ([]response.Orders, error) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}
	startIndex := (page - 1) * count
	endIndex := startIndex + count
	orderHistory, err := ou.orderRepo.UserOrderHistory(userID, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to get order history : %s", err)
	}

	return orderHistory, nil
}

func (ou *orderUseCase) GetOrderManagement(page, count int) (response.OrderManagement, error) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}
	startIndex := (page - 1) * count
	endIndex := startIndex + count
	var Datas response.OrderManagement
	orderHistory, err := ou.orderRepo.AllOrderData(startIndex, endIndex)
	if err != nil {
		return response.OrderManagement{}, fmt.Errorf("Failed to get order history : %s", err)
	}
	orderStatuses, err := ou.orderRepo.GetOrderStatuses()
	if err != nil {
		return response.OrderManagement{}, fmt.Errorf("Failed to retrieve order statuses : %s", err)
	}
	Datas.Orders = orderHistory
	fmt.Println("USECASE ORDERHISTORY ::", orderHistory)
	Datas.OrderStatuses = orderStatuses
	return Datas, nil
}

func (ou *orderUseCase) AllOrderOverView(page, count int) ([]response.Orders, error) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 10
	}
	startIndex := (page - 1) * count
	endIndex := startIndex + count

	AllOrders, err := ou.orderRepo.AllOrderData(startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to get order history : %s", err)
	}
	return AllOrders, nil
}

func (ou *orderUseCase) UpdateOrderStatus(statusID int, orderID int) error {
	updatedOrder, err := ou.orderRepo.ChangeOrderStatus(statusID, orderID)
	if err != nil {
		return fmt.Errorf("Failed to update order : %s", err)
	}
	if updatedOrder.ID == 0 {
		return fmt.Errorf("Failed to verify order status by id ")
	}

	return nil
}

func (ou *orderUseCase) OrderCancellation(orderID int) error {
	//find the order by provided orderID
	CancellingOrder, err := ou.orderRepo.FindOrderById(orderID)
	if err != nil {
		return fmt.Errorf("Failed to find order  :%s ", err)
	}

	if CancellingOrder.ID == 0 {
		return fmt.Errorf("Failed to veirfy order  by id")
	}
	PaymentMethodUsed, err := ou.paymentRepo.FindPaymentMethodById(CancellingOrder.PaymentMethodID)
	if err != nil {
		return fmt.Errorf("Failed to fetch payment method :%s", err)
	}

	//checking user has used coupon or not
	//reduct coupon percentage and insert money into wallet
	if CancellingOrder.CouponID != 0 && PaymentMethodUsed.MethodName == "online payment" {
		CouponUsedOrders, err := ou.orderRepo.FindOrdersUsedByCoupon(int(CancellingOrder.CouponID))
		if err != nil {
			return fmt.Errorf("Failed to fetch orders purchased using coupon : %s", err)
		}
		if len(CouponUsedOrders) == 0 {
			return fmt.Errorf("Failed to get coupon used orders")
		}
		var TotalPriceOfOrder float32
		for _, order := range CouponUsedOrders {
			TotalPriceOfOrder += order.Price
		}
		Coupon, err := ou.couponRepo.FindCouponById(int(CancellingOrder.CouponID))

		if err != nil {
			return fmt.Errorf("Failed to fetch coupon details :%s", err)
		}
		if Coupon.ID == 0 {
			return fmt.Errorf("Failed find coupon for reducting coupon percentage from the refunding amount")
		}
		TotalDiscountedAmountOnOrders := (TotalPriceOfOrder * float32(Coupon.DiscountPercent)) / 100
		DiscountAllowedPerOrder := TotalDiscountedAmountOnOrders / float32(len(CouponUsedOrders))

		Status, err := ou.orderRepo.GetStatusCancelled()
		if err != nil {
			return fmt.Errorf("Failed to procceed cancellation : %s", err)
		}
		if Status.ID == 0 {
			return fmt.Errorf("Failed to verify the status")
		}
		UpdatedOrder, err := ou.orderRepo.ChangeOrderStatus(int(Status.ID), orderID)
		if err != nil {
			return fmt.Errorf("Failed to update order status :%s", err)
		}
		if UpdatedOrder.ID == 0 {
			return fmt.Errorf("Failed to verify the orderline")
		}

		RefundingAmount := (CancellingOrder.Price - DiscountAllowedPerOrder)
		Wallet, err := ou.orderRepo.FindUserWallet(int(CancellingOrder.UserID))
		if err != nil {
			return fmt.Errorf("Failed to find user wallet : %s", err)
		}
		if Wallet.ID == 0 {
			NewWallet, err := ou.orderRepo.InitializeNewWallet(int(CancellingOrder.UserID))
			if err != nil {
				return fmt.Errorf("Failed to initialize wallet for user id %d", CancellingOrder.UserID)
			}
			if NewWallet.ID == 0 {
				return fmt.Errorf("Failed to verify initialized wallet")
			}
		}
		NewWalletBalance := (Wallet.Amount + RefundingAmount)
		UpdatedWalletDetails, err := ou.orderRepo.UpdateUserWallet(int(CancellingOrder.UserID), NewWalletBalance)
		if err != nil {
			return fmt.Errorf("order-cancelled failed to add amount to wallet : %s", err)
		}
		if UpdatedWalletDetails.ID == 0 {
			return fmt.Errorf("Failed to verify the wallet info by id ")
		}
	}

	if PaymentMethodUsed.MethodName == "cash on delivery" {
		Status, err := ou.orderRepo.GetStatusCancelled()
		if err != nil {
			return fmt.Errorf("Failed to procceed cancellation : %s", err)
		}
		if Status.ID == 0 {
			return fmt.Errorf("Failed to verify the status")
		}
		UpdatedOrder, err := ou.orderRepo.ChangeOrderStatus(int(Status.ID), orderID)
		if err != nil {
			return fmt.Errorf("Failed to update order status :%s", err)
		}
		if UpdatedOrder.ID == 0 {
			return fmt.Errorf("Failed to verify the orderline")
		}

	}
	return nil
}

func (ou *orderUseCase) GetUserWallet(userID int) (response.Wallet, error) {
	Wallet, err := ou.orderRepo.FindUserWallet(userID)
	if err != nil {
		return response.Wallet{}, fmt.Errorf("Failed to get user wallet :%s", err)
	}
	if Wallet.ID == 0 {
		return response.Wallet{}, fmt.Errorf("User dont have an wallet")
	}
	return Wallet, nil
}

func (ou *orderUseCase) CreateUserWallet(userID int) error {
	Wallet, err := ou.orderRepo.InitializeNewWallet(userID)
	if err != nil {
		return fmt.Errorf("Failed to initialize wallet for user %d : %s", userID, err)
	}
	if Wallet.ID == 0 {
		return fmt.Errorf("Failed to verify new wallet for user id  %d", userID)
	}
	return nil
}
