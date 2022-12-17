package main

import "fmt"

func main() {

	//Closure
	total := func() int {
		return soma(1, 34, 5, 80, 90, 9000) + 2
	}()

	fmt.Println(total)

	tot := func() int {
		return mult(5) + 5
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

func mult(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
