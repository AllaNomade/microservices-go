package main

func main() {

	a := "a"
	println(&a)

	var pointer *string = &a
	*pointer = "20"
	b := &a
	*b = "17"
	println(&b)
}
