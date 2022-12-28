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
	t := template.Must(template.New("CursoTemplate").Parse("Estudo, {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	// templa := template.New("CursoGotemplate")
	// templa, err := templa.Parse("Estudo, {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	// if err != nil {
	// 	panic(err)
	// }

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
