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

func NewRouter(router *gin.Engine, user *controller.UserController, transaction *controller.TransactionController, product *controller.ProductHandler, middleware *Middleware) *Router {
	return &Router{
		router:      router,
		user:        user,
		product:     product,
		transaction: transaction,
		middleware:  middleware,
	}
}

func (r *Router) Start(port string) {
	r.router.Use(r.middleware.Trace)

	// auth
	auth := r.router.Group("/auth")
	auth.POST("/register", r.user.Register)
	auth.POST("/login", r.user.Login)

	product := r.router.Group("/product")
	product.GET("", r.product.GetProducts)
	product.POST("", r.product.CreateProduct)
	product.GET("/id/:productId", r.product.FindProductByID)
	product.PUT("/id/:productId", r.product.UpdateProduct)
	product.DELETE("/id/:productId", r.product.DeleteProduct)

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

	users.POST("/admin", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminCreateEmployee, []string{"admin", "owner"}))
	users.GET("/admin", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminGetAllEmployee, []string{"admin", "owner"}))
	users.GET("/admin/:id", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminGetEmployeeById, []string{"admin", "owner"}))
	users.PUT("/admin/:id", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminUpdateEmployee, []string{"admin", "owner"}))
	users.DELETE("/admin/:id", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminDeleteEmployee, []string{"admin", "owner"}))

	// transaction
	transactions := r.router.Group("/transactions")
	transactions.POST("/inquire", r.transaction.Inquire)
	transactions.POST("confirm", r.middleware.Auth, r.middleware.CheckRole(r.transaction.Confirm, []string{"customer"}))
	transactions.GET("histories/me", r.middleware.Auth, r.transaction.FindAllByUserID)
	transactions.GET("histories/list", r.middleware.Auth, r.middleware.CheckRole(r.transaction.FindAll, []string{"admin", "kasir"}))
	transactions.PUT("id/:id/status", r.middleware.Auth, r.middleware.CheckRole(r.transaction.UpdateTrxStatus, []string{"kasir"}))

	err := r.router.Run(port)
	if err != nil {
		return
	}
}
