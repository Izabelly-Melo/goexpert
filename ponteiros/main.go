package main

func main() {
	// Memória -> Endereço -> Valor
	a := 10
	println(a)  // imprime o VALOR guardado na variável a
	println(&a) // imprime o ENDEREÇO de memória da variável a

	// cria um ponteiro que guarda o endereço de 'a'
	var ponteiro *int = &a

	// imprime o endereço que está dentro do ponteiro
	println(ponteiro)

	*ponteiro = 20 // acessa o endereço e muda o valor guardado nele

	println(*ponteiro) // valor acessado via ponteiro
	println(a)         // valor original foi alterado
}
