package main

func main() {

	a := 10
	defer println(&a)

	var pointer *int = &a
	*pointer = 20
	b := &a
	//&b, como na linha 15, aponta para o endereçamento da var b
	//caso passe println(b), retornará o endereçamento
	//na memória de a, linha 5, pois está apontando para a memória
	//como na linha 10.
	defer println(&b)
}
