package server

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router     *gin.Engine
	user       *controller.UserController
	product    *controller.ProductHandler
	middleware *Middleware
}

func NewRouter(router *gin.Engine, user *controller.UserController, product *controller.ProductHandler, middleware *Middleware) *Router {
	return &Router{
		router:     router,
		user:       user,
		product:    product,
		middleware: middleware,
	}
}

func (r *Router) Start(port string) {
	r.router.Use(r.middleware.Trace)

	auth := r.router.Group("/auth")
	auth.POST("/register", r.user.Register)
	auth.POST("/login", r.user.Login)

	users := r.router.Group("/users")
	users.POST("", r.middleware.Auth, r.user.CreateUser)
	users.GET("", r.middleware.Auth, r.middleware.CheckRole(r.user.GetUsers, []string{"admin", "owner"}))
	users.GET("/profile", r.middleware.Auth, r.user.Profile)
	users.GET("/email/:email", r.middleware.Auth, r.user.GetByEmail)
	users.PUT("/profile", r.middleware.Auth, r.user.UpdateUserProfile)

	product := r.router.Group("/product")
	product.GET("", r.product.GetProducts)
	product.POST("", r.product.CreateProduct)
	product.GET("/id/:productId", r.product.FindProductByID)
	product.PUT("/id/:productId", r.product.UpdateProduct)
	product.DELETE("/id/:productId", r.product.DeleteProduct)

	r.router.Run(port)
}
