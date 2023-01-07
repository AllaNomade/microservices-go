package main

import "fmt"

//map funções
func main() {
	salario := map[string]int{"Fulano": 100, "joao": 200}
	fmt.Println(salario["Fulano"])

	//deleta usuário
	// // delete(salario, "Fulano")

	//Add usuário
	salario["ciclano"] = 7000

	// //fmt.Println(salario)

	// Em map você escolhe o tipo da chave (string boolean ou int)
	salario1 := map[string]bool{"Terra Redonda": true, "Terra Plana": false}
	fmt.Println(salario1)

	//percorrer o Maps dos salários
	// // for nome, salarios := range salario {
	// // 	fmt.Printf("O salario de %s é %d\n", nome, salarios)
	// // }

	// // for _, salarios := range salario {
	// // 	fmt.Printf("O salario é %d\n", salarios)
	// // }

}
