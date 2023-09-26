package routes

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler, commonHandler *handler.AuthHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, referralHandler *handler.ReferralHandler) {

	router.POST("/su-login", adminHandler.SULogin)

	router.Use(middleware.AdminAuthJWT)
	{

		router.POST("/create-admin", middleware.AuthenticateSudoAdminJwt, middleware.Verified, adminHandler.CreateAdmin)

		category := router.Group("/category")
		{
			category.POST("/add-category", productHandler.CreateCategory)
			category.GET("/all-category", productHandler.ReadAllCategories)
			category.PUT("/update-category/:categoryID", productHandler.UpdateCategory)
			category.PUT("/block-category/:categoryID", productHandler.BlockCategory)
			category.PUT("/unblock-category/:categoryID", productHandler.UnBlockCategory)

		}

		products := router.Group("/products")
		{
			products.GET("/all-products", productHandler.DisplayAllProductsToAdmin)
			products.POST("/add-product/:categoryID", productHandler.CreateProduct)
			products.PUT("/update-product/:productID", productHandler.UpdateProduct)
			products.PUT("/block-product/:productID", productHandler.BlockProduct)
			products.PUT("/unblock-product/:productID", productHandler.UnBlockProduct)

		}

		coupon := router.Group("/promotions")
		{
			coupon.POST("/create-coupon", couponHandler.CreateCoupon)
			coupon.PUT("/update-coupon/:couponID", couponHandler.UpdateCoupon)
			coupon.GET("/all-coupons", couponHandler.ListOutAllCouponsToAdmin)
			coupon.PUT("/block-coupon/:couponID", couponHandler.BlockCoupon)
			coupon.PUT("/unblock-coupon/:couponID", couponHandler.UnBlockCoupon)
		}

		userManagement := router.Group("/user-management")
		{
			userManagement.GET("/view-all-users", adminHandler.DisplayAllUsers)
			userManagement.PUT("/block-user/:userID", adminHandler.BlockUser)
			userManagement.PUT("/unblock-user/:userID", adminHandler.UnblockUser)

		}

		orderManagement := router.Group("/orders")
		{
			orderManagement.GET("/", orderHandler.GetAllOrderOverViewPage)
			orderManagement.GET("/management", orderHandler.GetOrderManagementPage)
			orderManagement.PUT("/:orderID/update-status/:statusID", orderHandler.UpdateOrderStatus)
			router.GET("/sales-report", orderHandler.MonthlySalesReport)

		}

	}
}
