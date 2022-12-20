package main

import "fmt"

type myNumber int

//CONSTRAINTS
type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

//COMPARABLE compara a igualdade
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func main() {
	m := map[string]float64{"Fulano": 1.300, "joao": 2.300}
	m1 := map[string]int{"Fulano": 1300, "joao": 2300, "jess": 1500}

	m3 := map[string]int{"Fulano": 1300, "joao": 2300, "jess": 1500}

	fmt.Println(Soma(m))
	fmt.Println(Soma(m1))
	fmt.Println(Soma(m3))
	fmt.Println(Compara(10, 10))
}
