package database

import (
	"DB_LAB/internal/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.Inspector{},
		&entity.ShipOwner{},
		&entity.ShipType{},
		&entity.ShipSkipper{},
		&entity.Ship{},
		&entity.ShipOwnership{},
		&entity.TechnicalInspection{},
		&entity.Violation{},
		&entity.Query{},
	)

	if err != nil {
		panic(err)
	}
}
