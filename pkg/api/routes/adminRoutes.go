package routes

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler, authHandler *handler.AuthHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, referralHandler *handler.ReferralHandler, auth *middleware.AuthMiddleware) {

	router.POST("/login", authHandler.AdminLogin)

	router.Use(auth.AdminAuthRequired)
	{

		category := router.Group("/category")
		{

			category.POST("/add-category", productHandler.CreateCategory)
			category.POST("/add-image/:categoryID", productHandler.UploadCategoryImage)
			category.GET("/categories", productHandler.ReadAllCategories)
			category.PUT("/update-category/:categoryID", productHandler.UpdateCategory)
			category.PUT("/block-category/:categoryID", productHandler.BlockCategory)
			category.PUT("/unblock-category/:categoryID", productHandler.UnBlockCategory)

		}

		products := router.Group("/product")
		{
			products.POST("/add-product/:categoryID", productHandler.CreateProduct)
			products.POST("/add-images/:productID", productHandler.UploadProductImages)
			products.GET("/products", productHandler.ShowProductsToAdmin)
			products.GET("/all", productHandler.ShowProductsToAdmin)
			products.PUT("/update-product/:productID", productHandler.UpdateProduct)
			products.PUT("/block-product/:productID", productHandler.BlockProduct)
			products.PUT("/unblock-product/:productID", productHandler.UnBlockProduct)
			products.GET("/category/:categoryID", productHandler.ListProductsByCategoryAdmin)

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

		}
		router.GET("/sales-report", orderHandler.MonthlySalesReport)

	}
}
