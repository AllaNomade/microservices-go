package main

import "fmt"

func main() {
	exibealgo()
	nomes := map[string]int{"alan": 1200, "thiago": 1200, "filipe": 3000}

	fmt.Println(nomes)

	for nome, salario := range nomes {
		fmt.Printf("nome é: %s e seu salário é: %d\n", nome, salario)
	}
}

func exibealgo() {
	texto := []string{""}

	fmt.Println(texto)
}
