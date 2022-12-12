package main

import "fmt"

type ID int

var e float64 = 1.2
var f ID = 1

func main() {
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	for i, v := range array {
		fmt.Printf("o índice é %d e o valor é %d\n", i, v)
	}
}
