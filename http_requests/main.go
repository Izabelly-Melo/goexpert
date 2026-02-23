package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	// defer segura a excução da função para o final do escopo
	defer req.Body.Close() // fecha o corpo da resposta após a leitura

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resposta: %s\n", string(res))

}
