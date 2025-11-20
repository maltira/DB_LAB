package database

import (
	"DB_LAB/config"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		config.Env.DbHost,
		config.Env.DbUser,
		config.Env.DbPass,
		config.Env.DbName,
		config.Env.DbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().In(time.FixedZone("UTC+3", 3*60*60))
		},
	})

	if err != nil {
		panic(err)
	}
	return db
}
