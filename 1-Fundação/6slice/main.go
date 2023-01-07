package main

import "fmt"

// slice 1
func main() {
	s := []string{"alan", "thiago"}
	//len() mostra o numero de strings dentro do slice
	fmt.Println(len(s))
	//cap() mostra a capacidade total dentro do slice
	fmt.Println(cap(s))

	s = append(s, "alan")

	//len() mostra o numero de strings dentro do slice, neste caso, após o append() o numero de string no slice passou para 3
	fmt.Println(len(s))
	//cap() mostra a capacidade total dentro do slice, neste caso, após o append() a capacidade dobrou
	fmt.Println(cap(s))

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 6, 7)

	fmt.Println(len(a))
	fmt.Println(cap(a))

}
