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

type Order struct {
	ID       uint `gorm:"primaryKey"`
	Products []ProductOrder
}
type ProductOrder struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Product   Product
	OrderID   uint
	Order     Order
	Amount    int
}

type User struct {
	gorm.Model
	Username string
	Profile  StudentProfile
}

type StudentProfile struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	CompanyName string
	JobTile     string
	Level       string
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
		&User{},
		&StudentProfile{},
		&ProductOrder{},
	)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().AutoMigrate(
		&Product{},
		&Order{},
		&User{},
		&StudentProfile{},
		&ProductOrder{},
	)
	if err != nil {
		log.Fatal(err)
	}

	shirt := Product{
		Name:  "T-Shirt",
		Price: 350,
		Stock: 200,
	}

	short := Product{
		Name:  "T-Shirt V.2",
		Price: 400,
		Stock: 100,
	}

	toy := Product{
		Name:  "car toy",
		Price: 400,
		Stock: 100,
	}

	db.Create(&shirt)
	db.Create(&short)
	db.Create(&toy)

	order1 := Order{
		Products: []ProductOrder{
			{ProductID: shirt.ID, Amount: 1},
			{ProductID: short.ID, Amount: 1},
		},
	}

	db.Create(&order1)

	order2 := Order{
		Products: []ProductOrder{
			{ProductID: shirt.ID, Amount: 1},
			{ProductID: toy.ID, Amount: 1},
		},
	}

	db.Create(&order2)

}

//ถึง 1.58

// var found Product
// db.Preload("Orders").First(&found, 1)

// fmt.Printf("\n\n %+v \n\n", found)

// var found2 Order
// db.Preload("Product").First(&found2, 1)

// fmt.Printf("\n\n %+v \n\n", found2)

// user := User{
// 	Username: "pong",
// 	Profile: StudentProfile{
// 		CompanyName: "ODDS",
// 		JobTile:     "Golang Developer",
// 		Level:       "Poring",
// 	},
// }

// db.Save(&user)

// var foundUser User
// db.Preload("Profile").First(&foundUser, user.ID)

// fmt.Println(foundUser)
