package main

import (
	"github.com/Izabelly-Melo/goexpert/3/math"
	"github.com/google/uuid"
)

func main() {
	sum := math.NewMath(1, 3)
	println(sum.Sum())
	println(uuid.New().String())
}
