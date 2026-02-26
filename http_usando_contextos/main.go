package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()                          // cria um contexto vazio
	ctx, cancel := context.WithTimeout(ctx, time.Second) // cria um contexto com timeout de 1 segundo
	//ctx, cancel := context.WithCancel(ctx) // cria um contexto que pode ser cancelado manualmente
	defer cancel() // garante que o cancelamento do contexto seja chamado quando a função main terminar

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil) // cria uma requisição HTTP associada ao contexto
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
