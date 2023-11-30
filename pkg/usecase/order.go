package usecase

import (
	"errors"
	"fmt"
	"time"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

const (
	debit  = "debit"
	credit = "credit"
)

var (
	ErrNoOrders = errors.New("no orders created yet")
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

func (ou *orderUseCase) CheckOutDetails(userID int) (response.Checkout, error) {

	addresses, err := ou.userRepo.GetAllUserAddresses(userID)
	if err != nil {
		return response.Checkout{}, fmt.Errorf("Failed to retrieve checkout details %s", err)
	}

	if len(addresses) == 0 {
		return response.Checkout{}, fmt.Errorf("User don't have an address")
	}

	cartItems, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return response.Checkout{}, fmt.Errorf("Failed to retrieve cart items %s", err)
	}

	paymentMethods, err := ou.paymentRepo.GetPaymentMethods()
	if err != nil || paymentMethods == nil {
		return response.Checkout{}, fmt.Errorf("Failed to get payment methods %s", err)
	}

	return response.Checkout{
		Address:        addresses,
		Cart:           cartItems.Cart,
		Total:          cartItems.Total,
		Discount:       cartItems.Discount,
		PaymentOptions: paymentMethods,
	}, nil
}

func (ou *orderUseCase) GetRazorPayDetails(userID int) (response.PaymentDetails, error) {
	userCart, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to retrieve userCart :%s", err)
	}

	userData, err := ou.userRepo.FindUserByID(userID)
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
	address, err := ou.userRepo.FindDefaultAddress(userID)
	if err != nil || address.ID == 0 {
		return fmt.Errorf("Failed to find default address : %s", err)
	}

	addressID := address.ID

	cartData, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return fmt.Errorf("Failed to get Cart data :  %s", err)
	}

	couponDetails, err := ou.couponRepo.CheckAppliedCoupon(userID)
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
		Coupon, err := ou.couponRepo.FindCouponByID(couponDetails.CouponID)
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

		newOrderLine, err := ou.orderRepo.InsertOrder(request.NewOrder{

			UserID:          userID,
			ProductID:       int(productData.ProductID),
			AddressID:       int(addressID),
			Qty:             productData.Qty,
			Price:           productData.Price,
			PaymentMethodID: paymentMethodID,
			OrderStatusID:   int(statusID),
			CouponID:        couponDetails.CouponID,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		})

		if err != nil || newOrderLine.ID == 0 {
			return fmt.Errorf("Failed to insert order line : %s", err)
		}
	}

	err = ou.UpdateWallet(userID, cartData.Total, debit)
	if err != nil {
		return err
	}

	err = ou.cartUseCase.DeleteUserCart(userID)
	if err != nil {
		return fmt.Errorf("Failed to delete user cart :%s", err)
	}

	return nil
}

func (ou *orderUseCase) UpdateWallet(userID int, amount float32, transactionType string) error {
	wallet, err := ou.orderRepo.FindUserWalletByID(userID)
	if err != nil {
		return err
	}

	var updtAmount float32

	if transactionType == credit {
		updtAmount = (wallet.Amount + amount)

	}

	if transactionType == debit {
		updtAmount = (wallet.Amount - amount)

	}

	wallet, err = ou.orderRepo.UpdateUserWalletBalance(userID, updtAmount)
	if err != nil {
		return err
	}
	err = ou.UpdateWalletHistory(userID, amount, transactionType)
	if err != nil {
		return err
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

	orderHistory, err := ou.orderRepo.GetUserOrderHistory(userID, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to get order history : %s", err)
	}

	return orderHistory, nil
}

func (ou *orderUseCase) GetOrderManagement(page, count int) (response.OrderManagement, error) {
	startIndex, endIndex := helper.Paginate(page, count)

	orderHistory, err := ou.orderRepo.GetAllOrderData(startIndex, endIndex)
	if err != nil {
		return response.OrderManagement{}, fmt.Errorf("Failed to get order history : %s", err)
	}
	orderStatuses, err := ou.orderRepo.GetOrderStatuses()
	if err != nil {
		return response.OrderManagement{}, fmt.Errorf("Failed to retrieve order statuses : %s", err)
	}

	return response.OrderManagement{
		OrderStatuses: orderStatuses,
		Orders:        orderHistory,
	}, nil
}

func (ou *orderUseCase) AllOrderOverView(page, count int) ([]response.Orders, error) {
	startIndex, endIndex := helper.Paginate(page, count)

	allOrders, err := ou.orderRepo.GetAllOrderData(startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("Failed to get order history : %s", err)
	}
	return allOrders, nil
}

func (ou *orderUseCase) UpdateOrderStatus(statusID int, orderID int) error {
	updatedOrder, err := ou.orderRepo.ChangeOrderStatusByID(statusID, orderID)
	if err != nil {
		return fmt.Errorf("Failed to update order : %s", err)
	}
	if updatedOrder.ID == 0 {
		return fmt.Errorf("Failed to verify order status by id ")
	}

	return nil
}

func (ou *orderUseCase) ProcessReturnRequest(orderID int) error {
	order, err := ou.orderRepo.FindOrderByID(orderID)
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

	updatedOrder, err := ou.orderRepo.ChangeOrderStatusByID(int(status.ID), orderID)
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
	cancellingOrder, err := ou.orderRepo.FindOrderByID(orderID)
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

	orderStatus, err := ou.orderRepo.FindOrderStatusByID(cancellingOrder.OrderStatusID)
	if err != nil {
		return fmt.Errorf("Failed to find order statuses :%s", err)
	}

	//checking user has used coupon or not
	//reduce coupon percentage and insert money into wallet
	if cancellingOrder.CouponID != 0 && paymentMethodUsed.MethodName == "online payment" || paymentMethodUsed.MethodName == "Wallet" || orderStatus == "Returned" {
		couponClaimedOrders, err := ou.orderRepo.FindOrdersBoughtUsingCoupon(int(cancellingOrder.CouponID))
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

		coupon, err := ou.couponRepo.FindCouponByID(int(cancellingOrder.CouponID))

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
				return fmt.Errorf("Failed to proceed cancellation : %s", err)
			}
			if Status.ID == 0 {
				return fmt.Errorf("Failed to verify the status")
			}

			updatedOrder, err := ou.orderRepo.ChangeOrderStatusByID(int(Status.ID), orderID)
			if err != nil {
				return fmt.Errorf("Failed to update order status :%s", err)
			}
			if updatedOrder.ID == 0 {
				return fmt.Errorf("Failed to verify the order line")
			}

		}

		refundingAmount := (cancellingOrder.Price - discountedAmountPerOrderUsingCoupon)
		wallet, err := ou.orderRepo.FindUserWalletByID(int(cancellingOrder.UserID))
		if err != nil {
			return fmt.Errorf("Failed to find user wallet : %s", err)
		}
		if wallet.ID == 0 {
			newWallet, err := ou.orderRepo.InitializeNewUserWallet(int(cancellingOrder.UserID))
			if err != nil {
				return fmt.Errorf("Failed to initialize wallet for user id %d", cancellingOrder.UserID)
			}
			if newWallet.ID == 0 {
				return fmt.Errorf("Failed to verify initialized wallet")
			}
		}

		err = ou.UpdateWallet(wallet.UserID, refundingAmount, credit)
		if err != nil {
			return fmt.Errorf("Failed to update wallet %s:", err)
		}
	}

	if paymentMethodUsed.MethodName == "cash on delivery" && orderStatus != "Returned" {
		status, err := ou.orderRepo.GetStatusCancelled()
		if err != nil {
			return fmt.Errorf("Failed to proceed cancellation :%s", err)
		}
		if status.ID == 0 {
			return fmt.Errorf("Failed to verify the status")
		}

		updatedOrder, err := ou.orderRepo.ChangeOrderStatusByID(int(status.ID), orderID)
		if err != nil {
			return fmt.Errorf("Failed to update order status :%s", err)
		}
		if updatedOrder.ID == 0 {
			return fmt.Errorf("Failed to verify updated order")
		}
	}
	return nil
}

func (ou *orderUseCase) UpdateWalletHistory(userID int, amount float32, transactionType string) error {

	walletHistory, err := ou.orderRepo.UpdateWalletTransactionHistory(
		request.WalletTransactionHistory{
			TransactionTime: time.Now(),
			UserID:          userID,
			Amount:          amount,
			TransactionType: transactionType,
		},
	)

	if err != nil || walletHistory.ID == 0 {
		return func() error {
			if err != nil {
				return err
			}
			return fmt.Errorf("Failed to verify the updated history")
		}()

	}
	return nil
}

func (ou *orderUseCase) GetWalletHistory(userID int) ([]response.WalletTransactionHistory, error) {
	walletHistory, err := ou.orderRepo.GetWalletHistoryByUserID(userID)
	if err != nil {
		return walletHistory, err
	}

	return walletHistory, nil
}

func (ou *orderUseCase) GetUserWallet(userID int) (response.Wallet, error) {
	wallet, err := ou.orderRepo.FindUserWalletByID(userID)
	if err != nil {
		return response.Wallet{}, fmt.Errorf("Failed to get user wallet :%s", err)
	}
	if wallet.ID == 0 {
		return response.Wallet{}, fmt.Errorf("User don't have an wallet")
	}
	return wallet, nil
}

func (ou *orderUseCase) CreateUserWallet(userID int) error {

	wallet, err := ou.orderRepo.FindUserWalletByID(userID)
	if err != nil {
		return fmt.Errorf("Failed to check user wallet :%s", err)
	}

	if wallet.ID != 0 {
		return fmt.Errorf("User already have wallet")
	}

	newWallet, err := ou.orderRepo.InitializeNewUserWallet(userID)
	if err != nil {
		return fmt.Errorf("Failed to initialize wallet for user %d : %s", userID, err)
	}
	if newWallet.ID == 0 {
		return fmt.Errorf("Failed to verify new wallet for user id  %d", userID)
	}
	return nil
}

func (ou *orderUseCase) ValidateWalletPayment(userID int) error {
	wallet, err := ou.orderRepo.FindUserWalletByID(userID)
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

func (ou *orderUseCase) CreateInvoice(orderID int) (response.Invoice, error) {
	order, err := ou.orderRepo.FindOrderByID(orderID)
	if err != nil {
		return response.Invoice{}, fmt.Errorf("Failed to get order details:%s", err)
	}
	if order.ID == 0 {
		return response.Invoice{}, fmt.Errorf("Failed to fetch order by id")
	}

	invoiceData, err := ou.orderRepo.GetInvoiceDataByID(orderID)
	if err != nil {
		return response.Invoice{}, fmt.Errorf("Failed to get invoice data :%s", err)
	}
	if invoiceData.OrderID == 0 {
		return response.Invoice{}, fmt.Errorf("Failed to fetch order by id")
	}

	product, err := ou.productRepo.FindProductByID(invoiceData.ProductID)
	if err != nil {
		return response.Invoice{}, fmt.Errorf("Failed to get product data :%s", err)
	}
	if product.ID == 0 {
		return response.Invoice{}, fmt.Errorf("Failed to fetch product by id")
	}

	return response.Invoice{
		OrderDate:       order.CreatedAt.String(),
		OrderID:         orderID,
		DeliveryAddress: invoiceData.DeliveryAddress,
		ProductName:     product.ProductName,
		PaymentMethod:   invoiceData.PaymentMethod,
		ProductPrice:    product.Price,
		Discount:        (float32(product.Price) - order.Price),
		TotalAmount:     order.Price,
	}, nil
}

func (ou *orderUseCase) MonthlySalesReport() (response.MonthlySalesReport, error) {

	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()

	orders, err := ou.orderRepo.GetAllOrderData(0, 20)
	if err != nil {
		return response.MonthlySalesReport{}, fmt.Errorf("Failed to get order details :%s", err)
	}

	if len(orders) == 0 {
		return response.MonthlySalesReport{
			Date:           startDate.Format("January 2, 2006"),
			ReportFromDate: time.Now().Format("January 2, 2006"),
		}, ErrNoOrders
	}

	topSelling, err := ou.orderRepo.TopSellingProduct(startDate, endDate)
	if err != nil {
		return response.MonthlySalesReport{}, fmt.Errorf("Failed to find product :%s", err)
	}

	product, err := ou.productRepo.FindProductByID(topSelling.ProductID)
	if err != nil {
		return response.MonthlySalesReport{}, fmt.Errorf("Failed to find product :%s", err)
	}

	category, err := ou.productRepo.FindCategoryByID(product.CategoryID)
	if err != nil {
		return response.MonthlySalesReport{}, fmt.Errorf("Failed to find category :%s", err)
	}

	totalSalesCount, err := ou.orderRepo.GetTotalSaleCount(startDate, endDate)
	if err != nil {
		return response.MonthlySalesReport{}, fmt.Errorf("Failed to find total sales :%s", err)
	}

	avgOrderValue, err := ou.orderRepo.GetAverageOrderValue(startDate, endDate)
	if err != nil {
		return response.MonthlySalesReport{}, fmt.Errorf("Failed to find average order value :%s", err)
	}

	status, err := ou.orderRepo.GetStatusReturned()
	if err != nil {
		return response.MonthlySalesReport{}, fmt.Errorf("Failed to find returned products :%s", err)
	}

	revenueData, err := ou.orderRepo.GetTotalRevenue(int(status.ID), startDate, endDate)
	if err != nil {
		return response.MonthlySalesReport{}, fmt.Errorf("Failed to find total revenue :%s", err)
	}

	return response.MonthlySalesReport{

		Date:              time.Now().Format("January 2, 2006"),
		ReportFromDate:    startDate.Format("January 2, 2006"),
		TopSellingBrand:   category.Category_Name,
		TopSellingProduct: product.ProductName,
		TopSoldQuantity:   topSelling.Quantity,
		TotalSalesCount:   totalSalesCount,
		AverageOrderValue: avgOrderValue,
		TotalRevenue:      float32(helper.CalculateTotalRevenue(revenueData...)),
	}, nil

}
