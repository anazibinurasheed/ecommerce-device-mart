package api

import (
	"log"

	_ "github.com/anazibinurasheed/project-device-mart/api/docs"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/handler"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/middleware"
	"github.com/anazibinurasheed/project-device-mart/pkg/api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"  // swagger embed files
	swagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@title			Project Device Mart API
//	@version		1.0
//	@description	A e-Commerce API in Go using Gin framework
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Anaz Ibinu Rasheed
//	@contact.url	https://www.linkedin.com/in/anaz-ibinu-rasheed-a2b461253/
//	@contact.email	anazibinurasheed@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securitydefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization

//	@host		localhost:3000
//	@BasePath	/api/v1

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, commonHandler *handler.AuthHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, referralHandler *handler.ReferralHandler, auth *middleware.AuthMiddleware) *ServerHTTP {

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	router.LoadHTMLGlob("web/template/*.html")

	routes.UserRoutes(router.Group("/api/v1"), userHandler, adminHandler, productHandler, commonHandler, cartHandler, orderHandler, couponHandler, referralHandler, auth)

	routes.AdminRoutes(router.Group("/api/v1/admin"), userHandler, adminHandler, productHandler, commonHandler, cartHandler, orderHandler, couponHandler, referralHandler, auth)

	return &ServerHTTP{

		engine: router,
	}
}

func (s *ServerHTTP) Start(port string) {

	log.Fatal(s.engine.Run(":" + port))
}
