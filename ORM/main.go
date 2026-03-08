package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	//db.Create(&Product{Name: "Laptop", Price: 999.99})

	//products := []Product{
	//	{Name: "Smartphone", Price: 499.99},
	//	{Name: "Tablet", Price: 299.99},
	//}
	//db.Create(&products)

	p := Product{}
	//db.First(&p, 1) // SELECT * FROM products WHERE id = 1
	//fmt.Printf("Product: %+v\n", p)

	db.First(&p, "name = ?", "Tablet") // SELECT * FROM products WHERE name = 'Laptop'
	fmt.Printf("Product: %+v\n", p)

	fmt.Println("_________________")

	// select all products
	products := []Product{}
	db.Find(&products)
	for _, p := range products {
		fmt.Printf("Product: %+v\n", p)
	}
}
