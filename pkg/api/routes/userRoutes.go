package routes

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler, authHandler *handler.AuthHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, referralHandler *handler.ReferralHandler, auth *middleware.AuthMiddleware) {

	router.POST("/send-otp", authHandler.SendOTP)
	router.POST("/verify-otp", authHandler.VerifyOTP)
	router.POST("/sign-up", authHandler.UserSignUp)
	router.POST("/login", authHandler.UserLogin)
	router.POST("/logout", authHandler.Logout)
	router.POST("/webhook", orderHandler.WebhookHandler)

	// Authentication middleware
	router.Use(auth.UserAuthRequired)
	{
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

		referral := router.Group("/referral")
		{
			referral.GET("/get-code", referralHandler.GetReferralCode)
			referral.POST("/claim", referralHandler.ApplyReferralCode)
		}

		wallet := router.Group("/wallet")
		{
			wallet.GET("/", orderHandler.ViewUserWallet)
			wallet.POST("/create", orderHandler.CreateUserWallet)
			wallet.GET("/history", orderHandler.WalletTransactionHistory)
		}

		category := router.Group("/category")
		{
			category.GET("/all", productHandler.Categories)

		}

		product := router.Group("/product")
		{
			product.GET("/all", productHandler.DisplayAllProductsToUser)
			product.GET("/:productID", productHandler.ViewIndividualProduct)
			product.POST("/search", productHandler.SearchProducts)
			product.GET("/rating/:productID", productHandler.ValidateRatingRequest)
			product.POST("/rating/:productID", productHandler.AddProductRating)
			product.GET("/category/:categoryID", productHandler.ListProductsByCategoryUser)
		}

		wishlist := router.Group("/wishlist")
		{
			wishlist.GET("/", productHandler.ShowWishListProducts)
			wishlist.POST("/add/:productID", productHandler.AddToWishList)
			wishlist.DELETE("/remove/:productID", productHandler.RemoveFromWishList)
		}

		cart := router.Group("/cart")
		{
			cart.GET("/", cartHandler.ViewCart)
			cart.POST("/add/:productID", cartHandler.AddToCart)
			cart.PUT("/:productID/increment", cartHandler.IncrementQuantity)
			cart.PUT("/:productID/decrement", cartHandler.DecrementQuantity)
			cart.DELETE("/remove/:productID", cartHandler.RemoveFromCart)
		}

		coupon := router.Group("/coupon")
		{
			coupon.GET("/available", couponHandler.ListOutAvailableCouponsToUser)
			coupon.POST("/apply", couponHandler.ApplyCoupon)
			coupon.DELETE("/remove/:couponID", couponHandler.RemoveAppliedCoupon)
		}

		checkout := router.Group("/checkout")
		{
			checkout.GET("/", orderHandler.CheckOutPage)
		}

		payment := router.Group("/payment")
		{
			payment.GET("/online", orderHandler.GetOnlinePayment)
			payment.POST("/online/process", orderHandler.ProcessOnlinePayment)
			payment.POST("/wallet", orderHandler.PayUsingWallet)
			payment.POST("/cod-confirm", orderHandler.ConfirmCodDelivery)

		}

		orders := router.Group("/orders")
		{
			orders.GET("/", orderHandler.UserOrderHistory)
			orders.POST("/cancel/:orderID", orderHandler.CancelOrder)
			orders.POST("/return/:orderID", orderHandler.ReturnOrder)
			orders.GET("/invoice/:orderID", orderHandler.CreateInvoice)
		}

	}

}
