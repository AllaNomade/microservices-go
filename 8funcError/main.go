package main

import (
	"errors"
	"fmt"
)

func main() {

	valor, err := erro(30, 10)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(valor)

	fmt.Println(soma(1, 2))
	fmt.Println(soma1(1, 2))
}

func soma(a, b int) int {
	return a / b
	// retorna 0 pois é um inteiro (int)
}

func soma1(a, b float64) float64 {
	return a / b
	// retorna 0.5 pois traz o resultado decimal da soma utilizando o float64
}

func erro(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("=> A soma é maior que 50 > ")
	}

	return a + b, nil
}
