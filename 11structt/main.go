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

func (c Client) desativar() {
	c.active = false

	fmt.Printf("o cliente %s foi desativado", c.name)
}

func main() {

	fulano := Client{
		name:   "Fulano",
		age:    24,
		active: true,
	}

	fulano.desativar()
	fulano.active = false

}
