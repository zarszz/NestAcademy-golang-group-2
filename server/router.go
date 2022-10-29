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

	// transaction
	transactions := r.router.Group("/transactions")
	transactions.POST("/inquire", r.transaction.Inquire)

	err := r.router.Run(port)
	if err != nil {
		return
	}
}
