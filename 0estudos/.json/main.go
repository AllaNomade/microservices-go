package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int
	Balance int
}

func main() {

	account := Account{Number: 9, Balance: 900}
	// defer fmt.Println(account)

	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	jsonPure := []byte(`{"Number":2,"Balance":200}`)
	var accountX Account
	err = json.Unmarshal(jsonPure, &accountX)
	if err != nil {
		panic(err)
	}
	fmt.Println(accountX)
}
