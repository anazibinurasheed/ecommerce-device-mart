package api

import (
	"log"

	_ "github.com/anazibinurasheed/project-device-mart/cmd/api/docs"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"  // swagger embed files
	swagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type ServerHTTP struct {
	engine *gin.Engine
}

//	@title			Project Device Mart API
//	@version		1.0
//	@description	A e-Commerce API in Go using Gin framework
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Anaz Ibinu Rasheed
//	@contact.url	https://www.linkedin.com/in/anaz-ibinu-rasheed-a2b461253/
//	@contact.email	anazibinurasheed@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:3000
// @BasePath	/api/v1
func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler, commonHandler *handler.CommonHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, referralHandler *handler.ReferralHandler) *ServerHTTP {

	Engine := gin.New()
	Engine.LoadHTMLGlob("templates/*.html")

	Engine.Use(gin.Logger())
	Engine.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	user := Engine.Group("/api/v1")
	{
		user.POST("/send-otp", commonHandler.SendOTP)
		user.POST("/verify-otp", commonHandler.VerifyOTP)
		user.POST("/sign-up", middleware.Verified, userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)
		user.POST("/logout", commonHandler.Logout)
		user.POST("/webhook", orderHandler.WebhookHandler)

		user.Use(middleware.AuthenticateUserJwt)
		{

			profile := user.Group("/profile")
			{
				profile.GET("/", userHandler.Profile)
				profile.GET("/add-address", userHandler.GetAddAddressPage)
				profile.POST("/add-address", userHandler.AddAddress)
				profile.POST("/address-default/:addressID", userHandler.SetDefaultAddress)
				profile.PUT("/update-address/:addressID", userHandler.UpdateAddress)
				profile.GET("/addresses", userHandler.GetAllAdresses)
				profile.DELETE("/delete-address/:addressID", userHandler.DeleteAddress)
				profile.POST("/edit-username", userHandler.EditUserName)
				profile.POST("/verify-password", userHandler.ChangePasswordRequest)
				profile.POST("/change-password", userHandler.ChangePassword)
			}

			product := user.Group("/product")
			{
				product.GET("/", productHandler.DisplayAllProductsToUser)
				product.GET("/:productID", productHandler.ViewIndividualProduct)
				product.POST("/search", productHandler.SearchProducts)
				product.GET("/rating/:productID", productHandler.ValidateRatingRequest)
				product.POST("/rating/:productID", productHandler.AddProductRating)
				product.GET("/category/:categoryID", productHandler.ListProductsByCategory)

			}

			cart := user.Group("/cart")
			{
				cart.GET("/", cartHandler.ViewCart)
				cart.POST("/add/:productID", cartHandler.AddToCart)
				cart.PATCH("/:productID/increment", cartHandler.IncrementQuantity)
				cart.PATCH("/:productID/decrement", cartHandler.DecrementQuantity)
				cart.DELETE("/remove/:productID", cartHandler.RemoveFromCart)

			}

			coupon := user.Group("/coupon")
			{
				coupon.GET("/available", couponHandler.ListOutAvailableCouponsToUser)
				coupon.POST("/apply", couponHandler.ApplyCoupon)
				coupon.DELETE("/remove/:couponID", couponHandler.RemoveAppliedCoupon)
			}

			checkout := user.Group("/checkout")
			{
				checkout.GET("/", orderHandler.CheckOutPage)

			}

			payment := user.Group("/payment")
			{
				payment.POST("/order-cod-confirmed", orderHandler.ConfirmCodDelivery)
				payment.GET("/razorpay", orderHandler.MakePaymentRazorpay)
				payment.POST("/razorpay/process-order", orderHandler.ProccessRazorpayOrder)
				payment.POST("/wallet", orderHandler.WalletPayment)

			}

			order := user.Group("/my-orders")
			{
				order.GET("/", orderHandler.UserOrderHistory)
				order.POST("/cancel/:orderID", orderHandler.CancelOrder)
				order.POST("/return/:orderID", orderHandler.ReturnOrder)
				order.GET("/invoice/:orderID", orderHandler.DownloadInvoice)
			}

			referral := user.Group("/referral")
			{
				referral.GET("/get-code", referralHandler.GetReferralCode)
				referral.POST("/claim", referralHandler.ApplyReferralCode)
			}

			wallet := user.Group("/wallet")
			{
				wallet.GET("/", orderHandler.ViewUserWallet)
				wallet.POST("/create", orderHandler.CreateUserWallet)

			}

		}

	}

	admin := Engine.Group("api/v1/admin")
	{
		admin.POST("/su-login", adminHandler.SULogin)

		admin.Use(middleware.AdminAuthJWT)
		{

			admin.POST("/create-admin", middleware.AuthenticateSudoAdminJwt, middleware.Verified, adminHandler.CreateAdmin)

			category := admin.Group("/category")
			{
				category.POST("/add-category", productHandler.CreateCategory)
				category.GET("/all-category", productHandler.ReadAllCategories)
				category.PATCH("/update-category/:categoryID", productHandler.UpdateCategory)
				category.PATCH("/block-category/:categoryID", productHandler.BlockCategory)
				category.PATCH("/unblock-category/:categoryID", productHandler.UnBlockCategory)

			}

			products := admin.Group("/products")
			{
				products.GET("/all-products", productHandler.DisplayAllProductsToAdmin)
				products.POST("/add-product/:categoryID", productHandler.CreateProduct)
				products.PATCH("/update-product/:productID", productHandler.UpdateProduct)
				products.PATCH("/block-product/:productID", productHandler.BlockProduct)
				products.PATCH("/unblock-product/:productID", productHandler.UnBlockProduct)

			}

			coupon := admin.Group("/promotions")
			{
				coupon.POST("/create-coupon", couponHandler.CreateCoupon)
				coupon.PUT("/update-coupon/:couponID", couponHandler.UpdateCoupon)
				coupon.GET("/all-coupons", couponHandler.ListOutAllCouponsToAdmin)
				coupon.PUT("/block-coupon/:couponID", couponHandler.BlockCoupon)
				coupon.PUT("/unblock-coupon/:couponID", couponHandler.UnBlockCoupon)
			}

			user_management := admin.Group("/user-management")
			{
				user_management.GET("/view-all-users", adminHandler.DisplayAllUsers)
				user_management.PATCH("/block-user/:userID", adminHandler.BlockUser)
				user_management.PATCH("/unblock-user/:userID", adminHandler.UnblockUser)

			}

			order_management := admin.Group("/orders")
			{
				order_management.GET("/", orderHandler.GetAllOrderOverViewPage)
				order_management.GET("/management", orderHandler.GetOrderManagementPage)
				order_management.PUT("/:orderID/update-status/:statusID", orderHandler.UpdateOrderStatus)
			}

		}
	}
	return &ServerHTTP{engine: Engine}
}

func (s *ServerHTTP) Start() {
	log.Fatal(s.engine.Run(":3000"))
}
