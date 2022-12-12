package main

import "fmt"

func main() {

	//Closure
	total := func() int {
		return soma(1, 34, 5, 80, 90, 9000) * 2
	}()

	fmt.Println(total)
}

func soma(numeros ...int) int {

	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
