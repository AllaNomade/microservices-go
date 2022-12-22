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
	res, err := io.ReadAll(chamada.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	chamada.Body.Close()
}
