package main

import (
	"log"

	"github.com/zarszz/NestAcademy-golang-group-2/adaptor"
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

	rajaOngkirAdaptor := adaptor.NewRajaOngkirAdapter(config.RajaongkirBaseUrl)

	userDetailRepo := gorm_postgres.NewUserDetailGormRepository(db)
	userDetailSvc := service.NewUserDetailService(userDetailRepo, rajaOngkirAdaptor)

	userRepo := gorm_postgres.NewUserRepoGormPostgres(db)
	userSvc := service.NewServices(userRepo)
	userHandler := controller.NewUserController(userSvc, userDetailSvc)

	transactionSvc := service.TransactionServicesNew(config)
	transactionHandler := controller.NewTransactionController(transactionSvc)

	router := gin.Default()
	router.Use(gin.Logger())

	middleware := server.NewMiddleware(userSvc)

	app := server.NewRouter(router, userHandler, transactionHandler, middleware)

	app.Start(":4444")
}
