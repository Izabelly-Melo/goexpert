package main

import "fmt"

/*
- Por baixo dos panos, slices utilizam arrays como estrutura base
- Um slice possui três partes: ponteiro para o array subjacente, tamanho (len), capacidade (cap)

- O operador ":" cria um novo slice apontando para o mesmo array, apenas alterando a janela de visualização
- s[:n] pega os primeiros n elementos
- s[n:] ignora os primeiros n elementos
- Esses cortes não removem dados da memória, apenas mudam o intervalo acessado

- Quando o append ultrapassa a capacidade, o Go realoca um novo array maior e copia os dados, usando uma estratégia interna de crescimento
*/

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d capacity=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d capacity=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d capacity=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	fmt.Printf("len=%d capacity=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
	s = append(s, 30)
	fmt.Printf("len=%d capacity=%d %v\n", len(s), cap(s), s)
}
