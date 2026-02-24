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
	tmp := template.New("Cursor Template")

	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}} horas")

	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
