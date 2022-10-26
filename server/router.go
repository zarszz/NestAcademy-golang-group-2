package server

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router     *gin.Engine
	user       *controller.UserController
	middleware *Middleware
}

func NewRouter(router *gin.Engine, user *controller.UserController, middleware *Middleware) *Router {
	return &Router{
		router:     router,
		user:       user,
		middleware: middleware,
	}
}

func (r *Router) Start(port string) {
	r.router.Use(r.middleware.Trace)

	auth := r.router.Group("/auth")
	auth.POST("/register", r.user.Register)
	auth.POST("/login", r.user.Login)

	r.router.Run(port)
}
