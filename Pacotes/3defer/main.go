package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//pacote http.get faz a chamada do endereço
	chamada, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	//o comando defer atrasa a linha, o que será feita por último
	defer chamada.Body.Close()
	res, err := io.ReadAll(chamada.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
