package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	println("-------------------------")

	names := []string{"Iza", "Lino", "Jack"}
	for i, name := range names {
		println(i, name)
	}

	println("-------------------------")

	for _, name := range names {
		println(name)
	}

	println("-------------------------")

	for i := range names {
		println(i)
	}

	println("-------------------------")

	j := 0
	for j < 5 {
		println(j)
		j++
	}

	//for {
	//	println("Loop infinito")
	//}
}
