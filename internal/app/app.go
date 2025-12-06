package app

import (
	"DB_LAB/config"
	"DB_LAB/internal/middleware"
	"DB_LAB/internal/module"
	"DB_LAB/pkg/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	// ? Инициализация ENV
	config.InitEnv()

	// ? Инициализация DB
	db := database.InitDatabase()
	database.Migrate(db)

	// ? Инициализация роутера
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	apiGroup := router.Group("/api")

	// ? Инициализация модулей
	module.InitAuthModule(apiGroup, db)
	module.InitUserModule(apiGroup, db)
	module.InitShipModule(db, apiGroup)
	module.InitOwnerModule(db, apiGroup)
	module.InitViolationModule(db, apiGroup)
	module.InitInspectionModule(db, apiGroup)
	module.InitInspectorModule(db, apiGroup)
	module.InitOwnershipModule(db, apiGroup)
	module.InitSkipperModule(db, apiGroup)
	module.InitQueryModule(db, apiGroup)

	// ? Запуск сервера
	fmt.Printf("Запуск сервера по адресу: http://%s:%s/api\n", config.Env.AppHost, config.Env.AppPort)
	err := router.Run(":" + config.Env.AppPort)
	if err != nil {
		panic(fmt.Sprintf("Не удалось запустить сервер: %s", err))
	}
}
