package main

import "fmt"

type Adress struct {
	num   int
	city  string
	state string
}

type Client struct {
	name   string
	age    int
	active bool
	locale Adress
}

type Business struct {
	nome   string
	active bool
}

type Person interface {
	Desativar() //intaface pode haver métodos, como por ex: Desativar(x int) int
}

func (b Business) Desativar() {
	b.active = false

	fmt.Printf("A empresa %s ESTÁ desativada.", b.nome)
}

func (c Client) Desativar() {
	c.active = false

	fmt.Printf("o cliente %s foi desativado", c.name)
}

func desativacao(person Person) {
	person.Desativar()
}

func main() {

	fulano := Client{
		name:   "Fulano",
		age:    24,
		active: true,
	}

	HopiOutlet := Business{
		nome:   "Hopi Outlet",
		active: true,
	}

	desativacao(HopiOutlet)
	desativacao(fulano)

}
