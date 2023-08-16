package usecase

import (
	"fmt"
	"time"

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
	productRepo interfaces.ProductRepository
}

// some times error will invoke from here
func NewOrderUseCase(UserUseCase interfaces.UserRepository, CartUseCase services.CartUseCase, paymentUseCase interfaces.PaymentRepository, OrderUseCase interfaces.OrderRepository, CouponUseCase interfaces.CouponRepository, productUseCase interfaces.ProductRepository) services.OrderUseCase {
	return &orderUseCase{
		userRepo:    UserUseCase,
		cartUseCase: CartUseCase,
		paymentRepo: paymentUseCase,
		orderRepo:   OrderUseCase,
		couponRepo:  CouponUseCase,
		productRepo: productUseCase,
	}
}

//safe for concurrent use by multiple goroutines without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.

func (ou *orderUseCase) CheckOutDetails(userID int) (response.CheckOut, error) {
	var checkOutDetails response.CheckOut

	addresses, err := ou.userRepo.GetAllUserAddresses(userID)
	if err != nil {
		return response.CheckOut{}, fmt.Errorf("Failed to retrieve checkout details %s", err)
	}
	if len(addresses) == 0 {
		return response.CheckOut{}, fmt.Errorf("User dont have an address")
	}

	cartItems, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return response.CheckOut{}, fmt.Errorf("Failed to retrieve cartitems %s", err)
	}

	paymentMethods, err := ou.paymentRepo.GetPaymentMethods()
	if err != nil || paymentMethods == nil {
		return response.CheckOut{}, fmt.Errorf("Failed to get payment methods %s", err)
	}
	checkOutDetails.Address = addresses
	checkOutDetails.Cart = cartItems.Cart
	checkOutDetails.Total = cartItems.Total
	checkOutDetails.Discount = cartItems.Discount
	checkOutDetails.PaymentOptions = paymentMethods

	return checkOutDetails, nil
}

func (ou *orderUseCase) GetRazorPayDetails(userID int) (response.PaymentDetails, error) {
	userCart, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to retrieve userCart :%s", err)
	}

	userData, err := ou.userRepo.FindUserById(userID)
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to find user  %s", err)
	}

	razorPayOrderID, err := helper.MakeRazorPayPaymentId(int(userCart.Total * 100))
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to get razorpay id %s", err)
	}

	return response.PaymentDetails{
		Username:        userData.UserName,
		RazorPayOrderID: razorPayOrderID,
		Amount:          int(userCart.Total),
	}, nil

}
func (ou *orderUseCase) VerifyRazorPayPayment(signature string, razorpayOrderID string, paymentID string) error {
	err := helper.VerifyRazorPayPayment(signature, razorpayOrderID, paymentID)
	if err != nil {
		return fmt.Errorf("Failed payment not success : %s", err)
	}
	return nil
}

func (ou *orderUseCase) ConfirmedOrder(userID int, paymentMethodID int) error {

	address, err := ou.userRepo.FindDefaultAddressById(userID)
	if err != nil || address.ID == 0 {
		return fmt.Errorf("Failed to find default address : %s", err)
	}
	addressID := address.ID

	cartData, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return fmt.Errorf("Failed to get Cart data :  %s", err)
	}
	couponDetails, err := ou.couponRepo.FindAppliedCouponByUserId(userID)
	if err != nil {
		return fmt.Errorf("Failed to retrieve the coupon tracking details  ;%s", err)
	}

	if couponDetails.CouponID != 0 {
		UpdatedCouponTracking, err := ou.couponRepo.UpdateCouponUsage(userID)
		if err != nil {
			return fmt.Errorf("Failed to update coupon usage :%s", err)
		}
		if UpdatedCouponTracking.ID == 0 {
			return fmt.Errorf("Failed to verify inserted coupon record")
		}
	}
	if couponDetails.ID != 0 {
		Coupon, err := ou.couponRepo.FindCouponById(couponDetails.CouponID)
		if err != nil {
			return fmt.Errorf("Failed to find coupon by id : %s", err)
		}
		if !helper.IsCouponValid(Coupon.ValidTill) {
			return fmt.Errorf("Failed applied Coupon is expired")
		}
	}
	status, err := ou.orderRepo.GetStatusPending()
	if err != nil {
		return fmt.Errorf("Failed to get order status :%s", err)
	}
	statusID := status.ID

	for _, productData := range cartData.Cart {
		createdAt := time.Now()
		updatedAt := time.Now()
		newOrderLine, err := ou.orderRepo.InsertOrderLine(userID, int(productData.ProductID), int(addressID), productData.Qty, productData.Price, paymentMethodID, int(statusID), couponDetails.CouponID, createdAt, updatedAt)
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

	var orderManagementDetails response.OrderManagement

	orderHistory, err := ou.orderRepo.AllOrderData(startIndex, endIndex)
	if err != nil {
		return response.OrderManagement{}, fmt.Errorf("Failed to get order history : %s", err)
	}
	orderStatuses, err := ou.orderRepo.GetOrderStatuses()
	if err != nil {
		return response.OrderManagement{}, fmt.Errorf("Failed to retrieve order statuses : %s", err)
	}
	orderManagementDetails.Orders = orderHistory
	orderManagementDetails.OrderStatuses = orderStatuses
	return orderManagementDetails, nil
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

	allOrders, err := ou.orderRepo.AllOrderData(startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to get order history : %s", err)
	}
	return allOrders, nil
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

func (ou *orderUseCase) ProcessReturnRequest(orderID int) error {
	order, err := ou.orderRepo.FindOrderById(orderID)
	if err != nil {
		return fmt.Errorf("Failed to fetch order details: %s", err)
	}
	if order.ID == 0 {
		return fmt.Errorf("Failed to verify order by id")
	}

	if !helper.IsValidReturn(order.CreatedAt) {
		return fmt.Errorf("Failed, order return period is ended")
	}
	status, err := ou.orderRepo.GetStatusReturned()
	if err != nil {
		return fmt.Errorf("Failed to get return status :%s", err)
	}

	updatedOrder, err := ou.orderRepo.ChangeOrderStatus(int(status.ID), orderID)
	if err != nil {
		return fmt.Errorf("Failed to update order status to returned :%s", err)
	}
	if updatedOrder.ID == 0 {
		return fmt.Errorf("Failed to verify returned order")
	}
	err = ou.OrderCancellation(orderID)
	if err != nil {
		return fmt.Errorf("Failed to return order :%s", err)
	}
	return nil
}

func (ou *orderUseCase) OrderCancellation(orderID int) error {
	//find the order by provided orderID
	cancellingOrder, err := ou.orderRepo.FindOrderById(orderID)
	if err != nil {
		return fmt.Errorf("Failed to find order  :%s ", err)
	}
	//
	if cancellingOrder.ID == 0 {
		return fmt.Errorf("Failed to verify order by id")
	}
	paymentMethodUsed, err := ou.paymentRepo.FindPaymentMethodById(cancellingOrder.PaymentMethodID)
	if err != nil {
		return fmt.Errorf("Failed to fetch payment method :%s", err)
	}
	orderStatus, err := ou.orderRepo.FindOrderStatusById(cancellingOrder.OrderStatusId)
	if err != nil {
		return fmt.Errorf("Failed to find order statuses :%s", err)
	}

	//checking user has used coupon or not
	//reduct coupon percentage and insert money into wallet
	if cancellingOrder.CouponID != 0 && paymentMethodUsed.MethodName == "online payment" || paymentMethodUsed.MethodName == "Wallet" || orderStatus == "Returned" {
		couponClaimedOrders, err := ou.orderRepo.FindOrdersUsedByCoupon(int(cancellingOrder.CouponID))
		if err != nil {
			return fmt.Errorf("Failed to fetch orders purchased using coupon : %s", err)
		}
		if len(couponClaimedOrders) == 0 {
			return fmt.Errorf("Failed to get coupon used orders")
		}

		var totalPriceOfOrder float32
		for _, order := range couponClaimedOrders {
			totalPriceOfOrder += order.Price
		}

		coupon, err := ou.couponRepo.FindCouponById(int(cancellingOrder.CouponID))

		if err != nil {
			return fmt.Errorf("Failed to fetch coupon details :%s", err)
		}

		var discountedAmountPerOrderUsingCoupon float32
		if coupon.ID != 0 {

			TotalDiscountedAmountOnOrders := (totalPriceOfOrder * float32(coupon.DiscountPercent)) / 100
			discountedAmountPerOrderUsingCoupon = TotalDiscountedAmountOnOrders / float32(len(couponClaimedOrders))
		}

		if orderStatus != "Returned" {
			Status, err := ou.orderRepo.GetStatusCancelled()
			if err != nil {
				return fmt.Errorf("Failed to procceed cancellation : %s", err)
			}
			if Status.ID == 0 {
				return fmt.Errorf("Failed to verify the status")
			}

			updatedOrder, err := ou.orderRepo.ChangeOrderStatus(int(Status.ID), orderID)
			if err != nil {
				return fmt.Errorf("Failed to update order status :%s", err)
			}
			if updatedOrder.ID == 0 {
				return fmt.Errorf("Failed to verify the orderline")
			}

		}

		refundingAmount := (cancellingOrder.Price - discountedAmountPerOrderUsingCoupon)
		wallet, err := ou.orderRepo.FindUserWallet(int(cancellingOrder.UserID))
		if err != nil {
			return fmt.Errorf("Failed to find user wallet : %s", err)
		}
		if wallet.ID == 0 {
			newWallet, err := ou.orderRepo.InitializeNewWallet(int(cancellingOrder.UserID))
			if err != nil {
				return fmt.Errorf("Failed to initialize wallet for user id %d", cancellingOrder.UserID)
			}
			if newWallet.ID == 0 {
				return fmt.Errorf("Failed to verify initialized wallet")
			}
		}

		newWalletBalance := (wallet.Amount + refundingAmount)
		updatedWalletDetails, err := ou.orderRepo.UpdateUserWallet(int(cancellingOrder.UserID), newWalletBalance)
		if err != nil {
			if cancellingOrder.ID == 0 {
				return fmt.Errorf("Failed to verify cancelling product")
			}
			return fmt.Errorf("order-cancelled failed to add amount to wallet : %s", err)
		}

		if updatedWalletDetails.ID == 0 {
			return fmt.Errorf("Failed to verify the wallet info by id")
		}
	}

	if paymentMethodUsed.MethodName == "cash on delivery" && orderStatus != "Returned" {
		status, err := ou.orderRepo.GetStatusCancelled()
		if err != nil {
			return fmt.Errorf("Failed to procceed cancellation :%s", err)
		}
		if status.ID == 0 {
			return fmt.Errorf("Failed to verify the status")
		}

		updatedOrder, err := ou.orderRepo.ChangeOrderStatus(int(status.ID), orderID)
		if err != nil {
			return fmt.Errorf("Failed to update order status :%s", err)
		}
		if updatedOrder.ID == 0 {
			return fmt.Errorf("Failed to verify updated order")
		}
	}
	return nil
}

func (ou *orderUseCase) GetUserWallet(userID int) (response.Wallet, error) {
	wallet, err := ou.orderRepo.FindUserWallet(userID)
	if err != nil {
		return response.Wallet{}, fmt.Errorf("Failed to get user wallet :%s", err)
	}
	if wallet.ID == 0 {
		return response.Wallet{}, fmt.Errorf("User don't have an wallet")
	}
	return wallet, nil
}

func (ou *orderUseCase) CreateUserWallet(userID int) error {

	wallet, err := ou.orderRepo.FindUserWallet(userID)
	if err != nil {
		return fmt.Errorf("Failed to check user wallet :%s", err)
	}

	if wallet.ID != 0 {
		return fmt.Errorf("User already have wallet")
	}

	newWallet, err := ou.orderRepo.InitializeNewWallet(userID)
	if err != nil {
		return fmt.Errorf("Failed to initialize wallet for user %d : %s", userID, err)
	}
	if newWallet.ID == 0 {
		return fmt.Errorf("Failed to verify new wallet for user id  %d", userID)
	}
	return nil
}

func (ou *orderUseCase) ValidateWalletPayment(userID int) error {
	wallet, err := ou.orderRepo.FindUserWallet(userID)
	if err != nil {
		return fmt.Errorf("Failed to find user wallet : %s", err)
	}
	if wallet.ID == 0 {
		return fmt.Errorf("User don't have a wallet ")
	}

	userCart, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return fmt.Errorf("Failed to fetch user cart : %s", err)
	}
	if userCart.Total > wallet.Amount {
		return fmt.Errorf("Insufficient balance")
	}
	return nil
}

func (ou *orderUseCase) CreateInvoice(orderID int) ([]byte, error) {
	orderLineData, err := ou.orderRepo.FindOrderById(orderID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get invoice data :%s", err)
	}
	if orderLineData.ID == 0 {
		return nil, fmt.Errorf("Failed to fetch order by id")
	}

	order, err := ou.orderRepo.GetInvoiceData(orderID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get invoice data :%s", err)
	}
	if order.OrderID == 0 {
		return nil, fmt.Errorf("Failed to fetch order by id")
	}

	product, err := ou.productRepo.FindProductById(order.ProductID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get product data :%s", err)
	}
	if product.ID == 0 {
		return nil, fmt.Errorf("Failed to fetch product by id")
	}

	invoiceData := map[string]interface{}{

		"   ": "",

		"Order Date": orderLineData.CreatedAt.String(),
		"Order ID":   fmt.Sprint(orderID),

		" ": "",

		"Delivery Address": order.DeliveryAddress,

		"  ": "",

		"Product name":   product.ProductName,
		"Payment method": order.PaymentMethod,

		"Total Amount": fmt.Sprint(product.Price),
	}
	return helper.GenerateInvoicePDF(invoiceData), nil
}
