package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Izabelly-Melo/goexpert/api/internal/entity"
	sqlite "github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", float64(rand.Intn(350)))
	assert.Nil(t, err)

	repository := NewProduct(db)

	err = repository.Create(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", float64(rand.Intn(350)))
	assert.NoError(t, err)

	repository := NewProduct(db)

	err = repository.Create(product)
	assert.Nil(t, err)

	productFound, err := repository.FindByID(product.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", float64(rand.Intn(350)))
	assert.Nil(t, err)

	repository := NewProduct(db)

	err = repository.Create(product)
	assert.Nil(t, err)

	product.Name = "Product 1.1"
	product.Price = 450

	err = repository.Update(product)
	assert.Nil(t, err)

	productFound, err := repository.FindByID(product.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, "Product 1.1", productFound.Name)
	assert.Equal(t, 450, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", float64(rand.Intn(350)))
	assert.Nil(t, err)

	repository := NewProduct(db)

	err = repository.Create(product)
	assert.Nil(t, err)

	err = repository.Delete(product.ID.String())
	assert.Nil(t, err)

	_, err = repository.FindByID(product.ID.String())
	assert.NotNil(t, err)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), 100)
		assert.NoError(t, err)

		db.Create(product)
	}
	repository := NewProduct(db)

	products, err := repository.FindAll(1, 10, "asc")

	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = repository.FindAll(2, 10, "asc")

	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = repository.FindAll(3, 10, "asc")

	assert.Nil(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}
