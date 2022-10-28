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

	productRepo := gorm_postgres.NewProductRepoGormPostgres(db)
	productSvc := service.NewProductServices(productRepo)
	productHandler := controller.NewProductHandler(productSvc)

	transactionRepo := gorm_postgres.NewTransactionGormRepository(db)
	transactionSvc := service.NewTransactionService(*rajaOngkirAdaptor, userRepo, productRepo, transactionRepo)
	transactionHandler := controller.NewTransactionController(transactionSvc)

	router := gin.Default()
	router.Use(gin.Logger())

	middleware := server.NewMiddleware(userSvc)

	app := server.NewRouter(router, userHandler, productHandler, transactionHandler, middleware)

	app.Start(":" + config.Port)
}
