package main

// type assertion é a forma de extrair um valor concreto a partir de uma interface
func main() {
	var minhaVar interface{} = "Iza"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	println(res)
	println(ok)
}
