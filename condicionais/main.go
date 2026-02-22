package main

func main() {
	age := 14

	if age >= 18 {
		println("You are an adult")
	} else {
		println("You are a minor")
	}

	//switch case
	switch age {
	case 0, 1, 2, 3, 4, 5:
		println("You are a baby")
	case 6, 7, 8, 9, 10:
		println("You are a child")
	case 11, 12, 13, 14, 15, 16, 17:
		println("You are a teenager")
	default:
		println("You are an adult")
	}
}
