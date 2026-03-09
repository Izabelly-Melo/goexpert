package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product // relação um-para-muitos com Product (has many)
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int      // chave estrangeira para Category
	Category     Category // relação com Category
	SerialNumber SerialNumber
	gorm.Model   // adiciona os campos CreatedAt, UpdatedAt e DeletedAt
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int // chave estrangeira para Product
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{}) // cria a tabela products no banco de dados

	//db.Create(&Product{Name: "Laptop", Price: 999.99})
	/*
		p := Product{}
		db.First(&p, 1) // SELECT * FROM products WHERE id = 1
		p.Price = 899.99
		db.Save(&p) // UPDATE products SET price = 899.99 WHERE id = 1

		db.Delete(&p) // DELETE FROM products WHERE id = 1

	*/
	/*
		category := Category{Name: "Eletronicos"}
		db.Create(&category) // INSERT INTO categories (name) VALUES ('Eletrônicos')

		// INSERT INTO products (name, price, category_id) VALUES ('Smartphone', 499.99, 1)
		db.Create(&Product{Name: "Smartphone", Price: 499.99, CategoryID: category.ID})

		db.Create(&SerialNumber{Number: "SN654321", ProductID: 1}) // INSERT INTO serial_numbers (number, product_id) VALUES ('SN654321', 1)

		category = Category{Name: "Cozinha"}
		db.Create(&category) // INSERT INTO categories (name) VALUES ('Cozinha')

		// INSERT INTO products (name, price, category_id) VALUES ('Panela', 499.99, 1)
		db.Create(&Product{Name: "Panela", Price: 499.99, CategoryID: category.ID})

		db.Create(&SerialNumber{Number: "SN654321", ProductID: 2}) // INSERT INTO serial_numbers (number, product_id) VALUES ('SN654321', 2)
	*/
	categories := []Category{}
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error // Preload carrega os produtos e os números de série associados a cada categoria
	if err != nil {
		panic("failed to query categories")
	}

	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.Products {
			println(" -", product.Name, "- Serial Number:", product.SerialNumber.Number)
		}
	}

}
