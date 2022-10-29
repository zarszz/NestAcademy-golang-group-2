package server

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router      *gin.Engine
	user        *controller.UserController
	transaction *controller.TransactionController
	middleware  *Middleware
}

func NewRouter(router *gin.Engine, user *controller.UserController, transaction *controller.TransactionController, middleware *Middleware) *Router {
	return &Router{
		router:      router,
		user:        user,
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

	users := r.router.Group("/users")
	users.POST("", r.middleware.Auth, r.user.CreateUser)
	users.GET("", r.middleware.Auth, r.middleware.CheckRole(r.user.GetUsers, []string{"admin", "owner"}))
	users.GET("/profile", r.middleware.Auth, r.user.Profile)
	users.GET("/email/:email", r.middleware.Auth, r.user.GetByEmail)
	users.PUT("/profile", r.middleware.Auth, r.user.UpdateUserProfile)

	users.POST("/admin", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminCreateEmployee, []string{"admin", "owner"}))
	users.GET("/admin", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminGetAllEmployee, []string{"admin", "owner"}))
	users.GET("/admin/:id", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminGetEmployeeById, []string{"admin", "owner"}))
	users.PUT("/admin/:id", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminUpdateEmployee, []string{"admin", "owner"}))
	users.DELETE("/admin/:id", r.middleware.Auth, r.middleware.CheckRole(r.user.AdminDeleteEmployee, []string{"admin", "owner"}))

	// transaction
	transactions := r.router.Group("/transactions")
	transactions.POST("/inquire", r.transaction.Inquire)

	err := r.router.Run(port)
	if err != nil {
		return
	}
}
