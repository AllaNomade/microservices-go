package main

import "fmt"

func main() {

	//Closure traz uma função dentro de uma função, como na linha 8 e 14
	total := func() int {
		return soma(1, 1, 1, 1, 1, 9000) + 2
	}()

	fmt.Println(total)

	tot := func() int {
		return sum(5) + 5
	}()

	fmt.Println(tot)
}

func soma(numeros ...int) int {

	tot := 0

	for _, numero := range numeros {
		tot += numero
	}

	return tot
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
