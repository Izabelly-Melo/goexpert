package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{Nome: "Go para Iniciantes", CargaHoraria: 40},
		{Nome: "Go Intermediário", CargaHoraria: 60},
		{Nome: "Go Avançado", CargaHoraria: 80},
	})
	if err != nil {
		panic(err)
	}
}
