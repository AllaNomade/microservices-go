package main

import "fmt"

func main() {

	//fmt.println(soma())retorna a soma dos números
	fmt.Println(soma(1, 4, 7, 8, 14, 77))
}

//retorna soma como se fosse um spread, retornando todos os parâmetros
func soma(numeros ...int) int {

	total := 0

	for _, numero := range numeros {
		total += numero
	}
	//fmt.println(numeros) retorna o total de numeros em soma, linha 7
	fmt.Println(numeros)

	return total
}
