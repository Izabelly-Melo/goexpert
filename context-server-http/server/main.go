package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

// evitar ficar fazendo processamento desnecessário quando o cliente já cancelou a requisição, por exemplo,
// fechando a aba do navegador ou clicando em "stop" antes de a resposta ser enviada. 
// O contexto é propagado automaticamente para as funções chamadas dentro do handler, permitindo que elas também sejam canceladas se o cliente cancelar a requisição.
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciando processamento da requisição")
	defer log.Println("Finalizando processamento da requisição")
	select {
	case <-time.After(5 * time.Second):
		// imprime no comand line stdout
		log.Println("Request processada com sucesso")
		// imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
