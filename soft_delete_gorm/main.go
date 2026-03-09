package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // adiciona os campos CreatedAt, UpdatedAt e DeletedAt
}

// soft delete: marca o registro como deletado, mas não o remove do banco de dados
func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}) // cria a tabela products no banco de dados

	//db.Create(&Product{Name: "Laptop", Price: 999.99})

	p := Product{}
	db.First(&p, 1) // SELECT * FROM products WHERE id = 1
	p.Price = 899.99
	db.Save(&p) // UPDATE products SET price = 899.99 WHERE id = 1

	db.Delete(&p) // DELETE FROM products WHERE id = 1

}
