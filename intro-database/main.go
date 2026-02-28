package main

import "github.com/google/uuid"

type Products struct {
	ID    string
	Name  string
	Price float64
}

func NewProducts(name string, price float64) *Products {
	return &Products{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

}
