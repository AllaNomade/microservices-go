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

type business struct {
	nome string
}

type Person interface {
	desativar()
}

func (b business) desativar() {

}

func (c Client) desativar() {
	c.active = false

	fmt.Printf("o cliente %s foi desativado", c.name)
}

func desativacao(person Person) {
	person.desativar()
}

func main() {

	fulano := Client{
		name:   "Fulano",
		age:    24,
		active: true,
	}

	minhaEmpresa := business{
		nome: "algo",
	}

	desativacao(minhaEmpresa)

}
