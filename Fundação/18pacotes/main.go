package main

import (
	"fmt"
	"golang/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "fiat"}
	fmt.Println(carro.Andar())
	fmt.Printf("A marca do carro é %s:\n", carro)
	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
