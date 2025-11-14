package database

import (
	"DB_LAB/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Env.DbHost,
		config.Env.DbUser,
		config.Env.DbPass,
		config.Env.DbName,
		config.Env.DbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
