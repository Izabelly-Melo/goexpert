package main

import "fmt"

type test string
type number float32

type Tipo interface {
	~string | ~float32
}

func main() {
	m := map[string]int{"one": 1000, "two": 2500}
	println(Soma(m))
	m2 := map[string]float64{"one": 2000.0, "two": 3500.0}
	println(Soma(m2))

	v1 := map[string]test{"Word": "Hello"}
	v2 := map[string]number{"Pi": 3.14, "Euler": 2.71}
	println(Teste(v1))
	fmt.Printf("%.2f\n", Teste(v2))

	println(Compara(40, 30))
}

func Compara[T comparable](a, b T) bool {
	return a == b
}

func Soma[T int | float64](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func Teste[T Tipo](v map[string]T) T {
	var value T
	for _, v := range v {
		value += v
	}

	return value
}
