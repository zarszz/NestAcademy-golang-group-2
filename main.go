package main

import (
	"log"

	"github.com/zarszz/NestAcademy-golang-group-2/config"
	"github.com/zarszz/NestAcademy-golang-group-2/db"
	"github.com/zarszz/NestAcademy-golang-group-2/server"
	"github.com/zarszz/NestAcademy-golang-group-2/server/controller"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository/gorm_postgres"
	"github.com/zarszz/NestAcademy-golang-group-2/server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variable", err)
	}

	db, err := db.ConnectGormDB(config)
	if err != nil {
		panic(err)
	}

	userRepo := gorm_postgres.NewUserRepoGormPostgres(db)
	userSvc := service.NewServices(userRepo)
	userHandler := controller.NewUserController(userSvc)

	router := gin.Default()
	router.Use(gin.Logger())

	middleware := server.NewMiddleware(userSvc)

	app := server.NewRouter(router, userHandler, middleware)

	app.Start(":4444")
}
