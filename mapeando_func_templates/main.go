package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	templates := []string{"content.html", "header.html", "footer.html"}

	t := template.New("content.html").Funcs(template.FuncMap{
		"ToUpper": strings.ToUpper,
	})

	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{Nome: "Go para Iniciantes", CargaHoraria: 40},
		{Nome: "Go Intermediário", CargaHoraria: 60},
		{Nome: "Go Avançado", CargaHoraria: 80},
	})
	if err != nil {
		panic(err)
	}
}
