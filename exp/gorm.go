package main

import (
	"fmt"
	"log"
	"time"

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
	ID       uint `gorm:"primaryKey"`
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

type Course struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

type Class struct {
	ID        uint `gorm:"primaryKey"`
	CourseID  uint
	Course    Course
	TrainerID uint
	Trainer   User
	Start     time.Time
	End       time.Time
	Seats     int
	Students  []ClassStudent
}

type ClassStudent struct {
	ID        uint `gorm:"primaryKey"`
	ClassID   uint
	StudentID uint
	Student   User
}

func main() {
	// ต่อ db หรือก็คือต่อกับ table plus
	url := "host=localhost user=peagolang password=supersecret dbname=peagolang port=54329 sslmode=disable"
	//ใช lib gorm ของ go นำข้อมูลเข้าสู่ db
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
		&Course{},
		&Class{},
		&ClassStudent{},
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
		&Course{},
		&Class{},
		&ClassStudent{},
	)
	if err != nil {
		log.Fatal(err)
	}
	// **สร้างข้อมูลที่จะทดสอบ

	// สร้าง course 1 course
	tdd := Course{
		Name:        "TDD",
		Description: "TDD is fun!",
	}
	db.Create(&tdd)
	// สร้าง User
	pong := User{Username: "pong"}
	gap := User{Username: "gap"}
	kane := User{Username: "kane"}
	jua := User{Username: "jua"} // Trainer

	db.Create(&pong)
	db.Create(&gap)
	db.Create(&kane)
	db.Create(&jua)
	// สร้าง class หลังจากนั้นใส่ข้อมูลของคลาส
	class := Class{
		CourseID:  tdd.ID,
		TrainerID: jua.ID,
		Start:     time.Date(2023, 5, 10, 9, 0, 0, 0, time.Local),
		End:       time.Date(2023, 5, 12, 17, 0, 0, 0, time.Local),
		Seats:     10,
		Students: []ClassStudent{
			{StudentID: pong.ID},
			{StudentID: gap.ID},
		},
	}

	db.Save(&class)
	//**แสดงผลของ class ที่สร้างขึ้นมา
	var foundClass Class
	db.Preload("Course").Preload("Trainer").Preload("Students.Student").First(&foundClass, class.ID)

	fmt.Println("#ID: ", foundClass.ID)
	fmt.Println("Name: ", foundClass.Course.Name)
	fmt.Println("Description: ", foundClass.Course.Description)
	fmt.Println("\tBy: ", foundClass.Trainer.Username)
	fmt.Println("\tDate: ", foundClass.Start, foundClass.End)
	fmt.Println("Students: ")
	for _, student := range foundClass.Students {
		fmt.Println("\tName: ", student.Student.Username)
	}
	// 	Name:  "T-Shirt",
	// 	Price: 350,
	// 	Stock: 200,
	// }

	// short := Product{
	// 	Name:  "T-Shirt V.2",
	// 	Price: 400,
	// 	Stock: 100,
	// }

	// toy := Product{
	// 	Name:  "car toy",
	// 	Price: 400,
	// 	Stock: 100,
	// }

	// db.Create(&shirt)
	// db.Create(&short)
	// db.Create(&toy)

	// order1 := Order{
	// 	Products: []ProductOrder{
	// 		{ProductID: shirt.ID, Amount: 1},
	// 		{ProductID: short.ID, Amount: 1},
	// 	},
	// }

	// db.Create(&order1)

	// order2 := Order{
	// 	Products: []ProductOrder{
	// 		{ProductID: shirt.ID, Amount: 1},
	// 		{ProductID: toy.ID, Amount: 1},
	// 	},
	// }

	// db.Create(&order2)

	// var foundOrder Order
	// db.Preload("Products.Product").First(&foundOrder, order1.ID)
	// fmt.Printf("\n\n %+v \n\n", foundOrder)
	// PrintOrder(foundOrder)
}

// func PrintOrder(order Order) {
// 	fmt.Println()
// 	fmt.Printf("Order ID: %v\n", order.ID)
// 	fmt.Println("Products:")
// 	for _, p := range order.Products {
// 		fmt.Printf("\t%v\t%v\t%v\n", p.Product.Name, p.Product.Price, p.Amount)
// 	}
// }

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
