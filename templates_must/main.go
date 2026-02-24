package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{
		Nome:         "Go para Iniciantes",
		CargaHoraria: 40,
	}
	t := template.Must(template.New("Curso Template").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}} horas"))

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
