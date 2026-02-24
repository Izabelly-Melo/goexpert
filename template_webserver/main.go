package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{Nome: "Go para Iniciantes", CargaHoraria: 40},
			{Nome: "Go Intermediário", CargaHoraria: 60},
			{Nome: "Go Avançado", CargaHoraria: 80},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
