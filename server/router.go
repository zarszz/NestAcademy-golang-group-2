package server

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router      *gin.Engine
	user        *controller.UserController
	product     *controller.ProductHandler
	transaction *controller.TransactionController
	middleware  *Middleware
}

func NewRouter(router *gin.Engine, user *controller.UserController, product *controller.ProductHandler, transaction *controller.TransactionController, middleware *Middleware) *Router {
	return &Router{
		router:      router,
		user:        user,
		product:     product,
		middleware:  middleware,
		transaction: transaction,
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

	transaction := r.router.Group("/transaction")
	transaction.POST("inquiry", r.middleware.Auth, r.transaction.Inquiry)
	transaction.POST("confirm", r.middleware.Auth, r.middleware.CheckRole(r.transaction.Confirm, []string{"customer"}))
	transaction.GET("histories/me", r.middleware.Auth, r.transaction.FindAllByUserID)
	transaction.GET("histories/list", r.middleware.Auth, r.middleware.CheckRole(r.transaction.FindAll, []string{"admin", "kasir"}))
	transaction.PUT("id/:id/status", r.middleware.Auth, r.middleware.CheckRole(r.transaction.UpdateTrxStatus, []string{"kasir"}))

	r.router.Run(port)
}
