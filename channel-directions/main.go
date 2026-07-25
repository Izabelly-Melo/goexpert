package main

import "fmt"

/*
*
* Exemplo 5 de channel (DIRECTIONS)
*
* Criado como Bidirecional: Em main, make(chan string) cria um canal que pode enviar e receber dados.
* - chan<- string (Apenas Envio): Na função recebe, o operador <- à direita do chan restringe a função a apenas enviar dados (hello <- nome). Tentar ler do canal aqui gera erro de compilação.
* - <-chan string (Apenas Recepção): Na função ler, o operador <- à esquerda do chan restringe a função a apenas ler dados (<-data). Tentar enviar dados gera erro de compilação.
*
* Por que usar?
* - Segurança: O compilador bloqueia usos incorretos do canal.
* - Clareza: Deixa explícita a responsabilidade de cada função no código.
 */
// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}

func recebe(nome string, hello chan<- string) { // só vai receber informação
	hello <- nome
}

func ler(data <-chan string) { // só vai esvaziar o channel
	fmt.Println(<-data)
}
