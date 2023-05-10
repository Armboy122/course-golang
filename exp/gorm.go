package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
	Stock int
}

func main() {
	url := "host=localhost user=peagolang password=supersecret dbname=peagolang port=54329 sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().AutoMigrate(
		&Product{},
	)
	if err != nil {
		log.Fatal(err)
	}

	shirt := Product{
		Name:  "T-Shirt",
		Price: 350,
		Stock: 200,
	}

	db.Create(&shirt)
}
