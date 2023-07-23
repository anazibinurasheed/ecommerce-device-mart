package api

import (
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

//	@host		devicemart.store
//	@BasePath	/api/v1
func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler, commonHandler *handler.CommonHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, refferalHandler *handler.RefferalHandler) *ServerHTTP {

	Engine := gin.New()
	Engine.LoadHTMLGlob("templates/*.html") //  loading html for razorpay payment

	// Add the Gin Logger middleware.
	Engine.Use(gin.Logger())
	Engine.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	user := Engine.Group("/api/v1")
	{

		user.POST("/send-otp", commonHandler.SendOtpToPhone) //sign up otp
		user.POST("/verify-otp", commonHandler.OtpValidater) // otp for verify signup phone number
		user.POST("/sign-up", middleware.IsVerified, userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)
		// user.GET("/refresh_token", commonHandler.RefreshToken)
		user.POST("/logout", commonHandler.Logout)
		user.POST("/webhook", orderHandler.WebhookHandler)

		//changed route path //verirify login otp in this route
		//it will send otp if the credentials are verified
		user.Use(middleware.AuthenticateUserJwt)
		{

			user.GET("/products", productHandler.DisplayAllProductsToUser)
			user.GET("/product-item/:productID", productHandler.ViewProductItem)
			user.POST("/products/search", productHandler.SearchProducts)
			user.GET("/product/rating/:productID", productHandler.ValidateRatingRequest)
			user.POST("/product/rating/:productID", productHandler.AddProductRating)
			user.GET("/products-by-category/:categoryID", productHandler.ListProductsByCategory)
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
				payment.GET("/razorpay/", orderHandler.MakePaymentRazorpay)
				payment.POST("/razorpay/process-order", orderHandler.ProccessRazorpayOrder)
				payment.POST("/wallet", orderHandler.WalletPayment)

			}
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
				profile.POST("/change-password", userHandler.ChangePassword) //m.authchangepass

			}

			order := user.Group("/my-orders")
			{
				order.GET("/", orderHandler.UserOrderHistory)
				order.POST("/cancel/:orderID", orderHandler.CancelOrder)
				order.POST("/return/:orderID")

			}
			referral := user.Group("/referral")
			{
				referral.GET("/get-code", refferalHandler.GetRefferalCode)
				referral.POST("/claim", refferalHandler.ApplyRefferalCode)
			}

			wallet := user.Group("/wallet")
			{
				wallet.GET("/", orderHandler.ViewUserWallet)
				wallet.POST("/create", orderHandler.CreateUserWallet)

			}

		}
	}
	//
	admin := Engine.Group("api/v1/admin")
	{
		admin.POST("/sudo/login", adminHandler.SudoAdminLogin)
		// admin.POST("/phone", commonHandler.SendOtpToPhone)
		//got error
		admin.Use(middleware.AdminAuthJWT)
		{

			admin.POST("/create-admin", middleware.AuthenticateSudoAdminJwt, middleware.IsVerified, adminHandler.AdminSignup) //before using create-admin api should  ensure that provided phone number is valid by sending otp .

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
				order_management.PUT("/:orderId/update-status/:statusID", orderHandler.UpdateOrderStatus)
			}
		}
	}
	return &ServerHTTP{engine: Engine}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(":3000")
}
