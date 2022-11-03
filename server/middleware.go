package server

import (
	"log"
	"strings"
	"time"

	"github.com/zarszz/NestAcademy-golang-group-2/helper"
	"github.com/zarszz/NestAcademy-golang-group-2/server/controller"
	"github.com/zarszz/NestAcademy-golang-group-2/server/service"
	"github.com/zarszz/NestAcademy-golang-group-2/server/view"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	userSvc *service.UserServices
}

func NewMiddleware(userSvc *service.UserServices) *Middleware {
	return &Middleware{
		userSvc: userSvc,
	}
}

func (m *Middleware) Auth(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	tokenArr := strings.Split(bearerToken, "Bearer ")

	if len(tokenArr) != 2 {
		c.Set("ERROR", "no token")
		info := view.AdditionalInfoError{
			Message: "Invalid token",
		}
		payload := view.ErrUnauthorized(info, "UNAUTHORIZED")
		controller.WriteErrorJsonResponse(c, view.ErrUnauthorized(payload, "INVALID_TOKEN"))
		return
	}

	myTok, err := helper.VerifyToken(tokenArr[1])
	if err != nil {
		c.Set("ERROR", err.Error())
		info := view.AdditionalInfoError{
			Message: "Invalid token",
		}
		payload := view.ErrUnauthorized(info, "UNAUTHORIZED")
		controller.WriteErrorJsonResponse(c, view.ErrUnauthorized(payload, "INVALID_TOKEN"))
		return
	}

	c.Set("USER_ID", myTok.UserId)
	c.Set("USER_EMAIL", myTok.Email)

	c.Next()

}

func (m *Middleware) CheckRole(next gin.HandlerFunc, roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetString("USER_ID")
		user, _ := m.userSvc.FindByID(id)

		isExist := false

		for _, role := range roles {
			if role == user.Role {
				isExist = true
				break
			}
		}

		if !isExist {
			info := view.AdditionalInfoError{
				Message: "you dont have access for this resources",
			}
			payload := view.ErrUnauthorized(info, "GET_ALL_USERS_FAIL")
			controller.WriteErrorJsonResponse(ctx, payload)
			return
		}

		next(ctx)
	}
}

func (m *Middleware) Trace(c *gin.Context) {
	now := time.Now()
	log.Printf("Get request with method :%v Path :%v\n", c.Request.Method, c.Request.URL)
	c.Next()
	isError := c.GetString("ERROR")
	if isError != "" {
		log.Printf("get error when try to get all typicode :%v\n", isError)
	}
	log.Printf("Finised request with method :%v Path :%v\n", c.Request.Method, c.Request.URL)

	end := time.Since(now).Milliseconds()
	log.Println("response time:", end)
}
