package matematica

func Soma[T int | float64](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}
