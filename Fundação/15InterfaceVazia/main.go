package main

import "fmt"

type x interface{}

func main() {

	var x interface{} = 10
	var y interface{} = 10

	show(x)
	show(y)
}

func show(t interface{}) {
	fmt.Printf("o tipo da variavel é %T e o valor é %v\n", t, t)
}
