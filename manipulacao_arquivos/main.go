package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criar arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Escrever no arquivo
	//tamanho, err := f.WriteString("Olá, mundo!") // grava a string no arquivo
	tamanho, err := f.Write([]byte("Olá, mundo!")) // grava um slice de bytes no arquivo
	if err != nil {
		panic(err)
	}
	fmt.Printf("Criado com sucesso, tamanho %d bytes\n", tamanho)
	f.Close() // fecha o arquivo após a escrita

	// Ler arquivo
	//file, err := os.Open("arquivo.txt") // abre o arquivo para leitura
	file, err := os.ReadFile("arquivo.txt") // lê o conteúdo do arquivo e retorna um slice de bytes
	if err != nil {
		panic(err)
	}
	fmt.Printf("Conteúdo do arquivo: %s\n", string(file))

	fmt.Println("-------------")
	// Leitura de linha por linha
	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer arquivo.Close() // fecha o arquivo após a leitura

	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 3) // buffer de 3 bytes
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Printf("Lido: %s\n", string(buffer[:n]))
	}

	arquivo.Close()
	err = os.Remove("arquivo.txt") // remove o arquivo criado
	if err != nil {
		panic(err)
	}
}
