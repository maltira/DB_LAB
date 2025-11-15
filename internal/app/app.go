package app

import (
	"DB_LAB/config"
	"DB_LAB/pkg/database"
)

func Run() {
	// ? Инициализация ENV
	config.InitEnv()

	// ? Инициализация DB
	db := database.InitDatabase()
	database.Migrate(db)

}
