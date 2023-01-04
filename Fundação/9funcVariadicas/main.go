package main

import "fmt"

func main() {

	soma(1, 34, 5, 80, 90, 9000)
}

//retorna soma como se fosse um spread
func soma(numeros ...int) int {

	total := 0

	for _, numero := range numeros {
		total += numero
	}
	fmt.Println(numeros)

	return total
}
