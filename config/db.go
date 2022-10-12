package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	config := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=orders_by port=5432 sslmode=disable")

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// db.Debug().AutoMigrate(models.Item{}, models.Order{})
	return db
}
