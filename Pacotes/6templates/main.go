package main

import (
	"html/template"
	"os"
)

type Estudo struct {
	Nome         string
	CargaHoraria int
}

func main() {

	curso := Estudo{"go", 40}
	templa := template.New("CursoGotemplate")
	templa, _ = templa.Parse("Estudo, {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	// Execute possui método Writer
	err := templa.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
