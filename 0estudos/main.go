package main

import (
	"errors"
	"fmt"
)

func main() {
	exibealgo()
	valor, err := nada(5, 1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
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

func nada(a, b int) (int, error) {

	if a+b >= 7 {
		return 0, errors.New("Soma correta")
	} else {
		return a + b, nil
	}
}
