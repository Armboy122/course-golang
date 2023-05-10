package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Price  float64
	Stock  int
	Orders []Order
}

type Order struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Product   Product
	Amount    int
}

func main() {
	url := "host=localhost user=peagolang password=supersecret dbname=peagolang port=54329 sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(
		&Product{},
		&Order{},
	)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().AutoMigrate(
		&Product{},
		&Order{},
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

	shirt2 := Product{
		Name:  "T-Shirt V.2",
		Price: 400,
		Stock: 100,
	}
	db.Create(&shirt2)

	order1 := Order{
		ProductID: shirt.ID,
		Amount:    2,
	}

	db.Create(&order1)

	order2 := Order{
		ProductID: shirt.ID,
		Amount:    3,
	}
	db.Create(&order2)

	var found Product
	db.Preload("Orders").First(&found, 1)

	fmt.Printf("\n\n %+v \n\n", found)

	var found2 Order
	db.Preload("Product").First(&found2, 1)

	fmt.Printf("\n\n %+v \n\n", found2)

}
