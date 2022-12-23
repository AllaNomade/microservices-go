package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	//tag adicionadas
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	defer fmt.Println(conta)

	//json.Marshal transforma a requisição em JSON
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	//O encoder pode salvar em um arquivo, ou no terminal, por exemplo.
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	//O json.Unmarshal transforma o JSON na conta, o contrário do método json.Marshal()
	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	fmt.Println(contaX.Saldo)
}
