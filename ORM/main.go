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

	/*
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
	*/

	/*
		products := []Product{}
		db.Limit(2).Offset(2).Find(&products) // offset: pula os 2 primeiros registros, limit: limita a quantidade de registros retornados
		for _, p := range products {
			fmt.Printf("Product: %+v\n", p)
		}

	*/

	/*
		//where
		products := []Product{}
		db.Where("price > ?", 300).Find(&products) // SELECT * FROM products WHERE price > 300
		for _, p := range products {
			fmt.Printf("Product: %+v\n", p)
		}

		fmt.Println("_________________")
		db.Where("name LIKE ?", "%l%").Find(&products) // SELECT * FROM products WHERE name LIKE '%l%'
		for _, p := range products {
			fmt.Printf("Product: %+v\n", p)
		}
	*/

	p := Product{}
	db.First(&p, 1) // SELECT * FROM products WHERE id = 1
	p.Price = 899.99
	db.Save(&p) // UPDATE products SET price = 899.99 WHERE id = 1

	p2 := Product{}
	db.First(&p2, 1) // SELECT * FROM products WHERE id = 2
	fmt.Printf("Product: %+v\n", p2)

	db.Delete(&p2) // DELETE FROM products WHERE id = 2
}
