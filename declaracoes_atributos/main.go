package main

const a = "Hello, World!"

// global scope
var (
	b bool
	c int
	d string
	e float64
)

func main() {
	testVariable(b, c, d, e)

	// local variable
	var f string
	println(f)
	println("--------------------------------")

	// declare variable with initial value
	var g = true
	println(g)
	println("--------------------------------")

	x := 100
	println(x)
	println("--------------------------------")
}

func testVariable(b bool, c int, d string, e float64) {
	println("--------------------------------")
	print("without setting value in boolean variable: ", b, "\n")
	print("setting value in boolean variable: ", varBool(), "\n")
	println("--------------------------------")
	print("without setting value in integer variable: ", c, "\n")
	print("setting value in integer variable: ", varInt(), "\n")
	println("--------------------------------")
	print("without setting value in string variable: ", d, "\n")
	print("setting value in string variable: ", varString(), "\n")
	println("--------------------------------")
	print("without setting value in float64 variable: ", e, "\n")
	print("setting value in float64 variable: ", varFloat64(), "\n")
	println("--------------------------------")
}

func varBool() bool {
	b = true
	return b
}

func varInt() int {
	c = 10
	return c
}

func varString() string {
	d = "Hello"
	return d
}

func varFloat64() float64 {
	e = 10.0
	return e
}
