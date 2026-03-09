package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := db.Begin() // inicia uma transação
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error // Select ... for update - nesse mesmo id, ficará pausado (em espera) até que execute o tx.Commit()
	if err != nil {
		panic(err)
	}

	c.Name = "Eletronicos"
	tx.Debug().Save(&c)
	tx.Commit()
}
