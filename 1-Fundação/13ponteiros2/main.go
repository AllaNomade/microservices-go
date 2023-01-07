package main

func soma(a, b *int) int {
	//memória do PONTEIRO
	*a = 5
	*b = 30
	return *a + *b
}

func main() {
	//CÓPIA do PONTEIRO
	minhaVar := 10
	minhaVar1 := 20
	soma(&minhaVar, &minhaVar1)
	println(minhaVar, minhaVar1)
}
