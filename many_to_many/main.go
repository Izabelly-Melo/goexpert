package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"` // relação muitos-para-muitos com Product
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Category   []Category `gorm:"many2many:product_categories;"` // relação muitos-para-muitos com Category
	gorm.Model            // adiciona os campos CreatedAt, UpdatedAt e DeletedAt
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{}) // cria a tabela products no banco de dados
	/*
		category := Category{Name: "Cozinha"}
		db.Create(&category) // INSERT INTO categories (name) VALUES ('Cozinha')
		category2 := Category{Name: "Eletronicos"}
		db.Create(&category2) // INSERT INTO categories (name) VALUES ('Eletrônicos')

		// INSERT INTO products (name, price, category_id) VALUES ('Panela', 499.99, 1)
		db.Create(&Product{Name: "Panela", Price: 499.99, Category: []Category{category, category2}}) // associa a panela com as categorias de cozinha e eletrônicos
	*/

	categories := []Category{}
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error // Preload carrega os produtos e os números de série associados a cada categoria
	if err != nil {
		panic("failed to query categories")
	}

	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.Products {
			println(" -", product.Name)
		}
	}
}
