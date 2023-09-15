package routes

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler, commonHandler *handler.CommonHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, referralHandler *handler.ReferralHandler) {

	// Define routes for various endpoints
	router.POST("/send-otp", commonHandler.SendOTP)
	router.POST("/verify-otp", commonHandler.VerifyOTP)
	router.POST("/sign-up", middleware.Verified, userHandler.UserSignUp)
	router.POST("/login", userHandler.UserLogin)
	router.POST("/logout", commonHandler.Logout)
	router.POST("/webhook", orderHandler.WebhookHandler)

	// Authentication middleware
	router.Use(middleware.AuthenticateUserJwt)
	{
		// Group routes for the "profile" section
		profile := router.Group("/profile")
		{
			profile.GET("/", userHandler.Profile)
			profile.GET("/add-address", userHandler.GetAddAddressPage)
			profile.POST("/add-address", userHandler.AddAddress)
			profile.POST("/address-default/:addressID", userHandler.SetDefaultAddress)
			profile.PUT("/update-address/:addressID", userHandler.UpdateAddress)
			profile.GET("/addresses", userHandler.GetAllAddresses)
			profile.DELETE("/delete-address/:addressID", userHandler.DeleteAddress)
			profile.POST("/edit-username", userHandler.EditUserName)
			profile.POST("/verify-password", userHandler.ChangePasswordRequest)
			profile.POST("/change-password", userHandler.ChangePassword)
		}

		// Group routes for the "referral" section
		referral := router.Group("/referral")
		{
			referral.GET("/get-code", referralHandler.GetReferralCode)
			referral.POST("/claim", referralHandler.ApplyReferralCode)
		}

		// Group routes for the "wallet" section
		wallet := router.Group("/wallet")
		{
			wallet.GET("/", orderHandler.ViewUserWallet)
			wallet.POST("/create", orderHandler.CreateUserWallet)
		}

		// Group routes for the "product" section
		product := router.Group("/product")
		{
			product.GET("/", productHandler.DisplayAllProductsToUser)
			product.GET("/:productID", productHandler.ViewIndividualProduct)
			product.POST("/search", productHandler.SearchProducts)
			product.GET("/rating/:productID", productHandler.ValidateRatingRequest)
			product.POST("/rating/:productID", productHandler.AddProductRating)
			product.GET("/category/:categoryID", productHandler.ListProductsByCategory)
		}

		// Group routes for the "cart" section
		cart := router.Group("/cart")
		{
			cart.GET("/", cartHandler.ViewCart)
			cart.POST("/add/:productID", cartHandler.AddToCart)
			cart.PUT("/:productID/increment", cartHandler.IncrementQuantity)
			cart.PUT("/:productID/decrement", cartHandler.DecrementQuantity)
			cart.DELETE("/remove/:productID", cartHandler.RemoveFromCart)
		}

		// Group routes for the "coupon" section
		coupon := router.Group("/coupon")
		{
			coupon.GET("/available", couponHandler.ListOutAvailableCouponsToUser)
			coupon.POST("/apply", couponHandler.ApplyCoupon)
			coupon.DELETE("/remove/:couponID", couponHandler.RemoveAppliedCoupon)
		}

		// Group routes for the "checkout" section
		checkout := router.Group("/checkout")
		{
			checkout.GET("/", orderHandler.CheckOutPage)
		}

		// Group routes for the "payment" section
		payment := router.Group("/payment")
		{
			payment.GET("/online", orderHandler.GetOnlinePayment)
			payment.POST("/online/process", orderHandler.ProcessOnlinePayment)
			payment.POST("/wallet", orderHandler.PayUsingWallet)
			payment.POST("/cod-confirm", orderHandler.ConfirmCodDelivery)

		}

		// Group routes for the "my-orders" section
		orders := router.Group("/orders")
		{
			orders.GET("/", orderHandler.UserOrderHistory)
			orders.POST("/cancel/:orderID", orderHandler.CancelOrder)
			orders.POST("/return/:orderID", orderHandler.ReturnOrder)
			orders.GET("/invoice/:orderID", orderHandler.CreateInvoice)
		}

	}

}
