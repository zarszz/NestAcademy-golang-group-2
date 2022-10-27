package db

import (
	"fmt"

	"github.com/zarszz/NestAcademy-golang-group-2/config"
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGormDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUsername, config.DbPassword, config.DbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbs, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = dbs.Ping()
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(model.User{}, model.UserDetail{})

	return db, nil
}
