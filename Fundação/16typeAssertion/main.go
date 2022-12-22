package main

import "fmt"

func main() {
	var minhaVar interface{} = "Alan Paiva"
	println(minhaVar.(string))

	res, ok := minhaVar.(int)

	fmt.Printf("o valor de res é %v e o OK é %v", res, ok)
}
