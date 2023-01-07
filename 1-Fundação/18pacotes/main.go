package main

import (
	"fmt"
	"gopackage/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "fiat"}
	fmt.Println(carro.Andar())
	fmt.Printf("A marca do carro é %s:\n", carro)
	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(matematica.A)
}
